package file

import (
    "io/ioutil"
    "log"
    "path/filepath"

    "github.com/fsnotify/fsnotify"
)

var (
    // ConfigTypeJSON ...
    ConfigTypeJSON ConfigType = "json"
    // ConfigTypeToml ...
    ConfigTypeToml ConfigType = "toml"
    // ConfigTypeYaml ...
    ConfigTypeYaml ConfigType = "yaml"
)

type FileHandle struct {
    path        string
    enableWatch bool
    changed     chan struct{}
}

func init() {
}

func extParser(cfgPath string) ConfigType {
    ext := filepath.Ext(cfgPath)
    switch ext {
    case ".json":
        return ConfigTypeJSON
    case ".toml":
        return ConfigTypeToml
    case ".yaml":
        return ConfigTypeYaml
    default:
        log.Panic("data source: invalid configuration type")
    }
    return ""
}

func (fp *FileHandle) Parse(path string, watch bool) ConfigType {
    absolutePath, err := filepath.Abs(path)
    if err != nil {
        log.Panic("new datasource", elog.FieldErr(err))
    }
    fp.path = absolutePath
    fp.enableWatch = watch

    if watch {
        fp.changed = make(chan struct{}, 1)
        go fp.watch()
    }

    return extParser(path)
}

func (fp *FileHandle) ReadConfig() (content []byte, err error) {
    return ioutil.ReadFile(fp.path)
}

func (fp *FileHandle) Close() error {
    close(fp.changed)
    return nil
}

func (fp *FileHandle) IsConfigChanged() <-chan struct{} {
    return fp.changed
}

// Watch file and automate update.
func (fp *FileHandle) watch() {
    w, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal("new file watcher", err)
    }
    defer w.Close()

    configFile := filepath.Clean(fp.path)
    realConfigFile, _ := filepath.EvalSymlinks(fp.path)

    log.Info("read watch",
        "configFile", configFile,
        "realConfigFile", realConfigFile,
        "fppath", fp.path,
    )

    done := make(chan bool)
    go func() {
        for {
            select {
            case event := <-w.Events:
                currentConfigFile, _ := filepath.EvalSymlinks(fp.path)

                log.Info("read watch event",
                    "event", filepath.Clean(event.Name),
                    "path", filepath.Clean(fp.path),
                    "currentConfigFile", currentConfigFile,
                    "realConfigFile", realConfigFile,
                )
                // we only care about the config file with the following cases:
                // 1 - if the config file was modified or created
                // 2 - if the real path to the config file changed (eg: k8s ConfigMap replacement)
                const writeOrCreateMask = fsnotify.Write | fsnotify.Create
                if (filepath.Clean(event.Name) == configFile &&
                    event.Op&writeOrCreateMask != 0) ||
                    (currentConfigFile != "" && currentConfigFile != realConfigFile) {

                    realConfigFile = currentConfigFile
                    log.Info("modified file", event.Name, realConfigFile)

                    select {
                    case fp.changed <- struct{}{}:
                    default:

                    }
                }
            case err := <-w.Errors:
                log.Error("read watch error", err)
            }
        }
    }()
    err = w.Add(fp.path)
    if err != nil {
        log.Fatal(err)
    }
    <-done
}
