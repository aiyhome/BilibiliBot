package file

import (
    "encoding/json"
    "errors"
    "github.com/BurntSushi/toml"
    "gopkg.in/yaml.v3"
    "net/url"
)

// ConfigType 配置类型
type ConfigType string

// 文件解析器
type Unmarshaller = func([]byte, interface{}) error

type IDataHandle interface {
    Parse(addr string, watch bool) ConfigType
    ReadConfig() ([]byte, error)
    IsConfigChanged() <-chan struct{}
    io.Closer
}

var (
    // ErrInvalidDataSource defines an error that the scheme has been registered
    ErrInvalidDataSource = errors.New("invalid data source, please make sure the scheme has been registered")
    // ErrInvalidUnmarshaller ...
    ErrInvalidUnmarshaller = errors.New("invalid unmarshaller, please make sure the config type is right")
    // ErrDefaultConfigNotExist 默认配置不存在
    ErrDefaultConfigNotExist = errors.New("default config not exist")

    registry map[string]FileHandle

    // 文件协议
    FileScheme = "file"

    unmarshallerMap = map[ConfigType]Unmarshaller{
        ConfigTypeJSON: json.Unmarshal,
        ConfigTypeToml: toml.Unmarshal,
        ConfigTypeYaml: yaml.Unmarshal,
    }
)

func init() {
    registry = make(map[string]IDataHandle)
    Regist(FileScheme, &FileHandle{})
}

// Register registers a dataSource handler function to the registry
func Regist(scheme string, handler IDataHandle) {
    registry[scheme] = handler
}

func New(cfgPath string, watch bool) (IDataHandle, Unmarshaller, ConfigType, error) {
    scheme := FileScheme
    urlObj, err := url.Parse(cfgPath)
    if err == nil && len(urlObj.Scheme) > 1 {
        scheme = urlObj.Scheme
    }

    // 如果是默认file协议，判断下文件是否存在
    if scheme == FileScheme {
        _, err := os.Stat(cfgPath)
        if err != nil {
            return nil, nil, "", ErrDefaultConfigNotExist
        }
    }

    handleFunc, exist := registry[scheme]
    if !exist {
        return nil, nil, "", ErrInvalidDataSource
    }
    tag := handleFunc.Parse(cfgPath, watch)

    parser, flag := unmarshallerMap[tag]
    if !flag {
        return nil, nil, "", ErrInvalidUnmarshaller
    }
    return handleFunc, parser, tag, nil
}
