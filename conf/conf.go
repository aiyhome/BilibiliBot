package conf

import (
	"BilibiliBot/util/log"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"strings"
)

type Object interface{}

type ConfigFileData struct {
	Path  string
	Obj   Object
	Cmd   string
	Usage string
}

var (
	Common    = &CommonConf{}
	Platforms = &PlatformsConf{}
	Users     = &UsersConf{}
	ConfLst   = []ConfigFileData{
		{
			Path:  "./toml/config.toml",
			Obj:   Common,
			Cmd:   "c",
			Usage: "common config file path",
		},
		{
			Path:  "./toml/platforms.toml",
			Obj:   Platforms,
			Cmd:   "p",
			Usage: "platforms config file path",
		},
		{
			Path:  "./toml/users.toml",
			Obj:   Users,
			Cmd:   "u",
			Usage: "users config file path",
		},
	}

	TestConf = &CommonConf{}
)

func init() {
	for _, v := range ConfLst {
		flag.StringVar(&v.Path, v.Cmd, v.Path, v.Usage)
	}
}

func Load() (err error) {
	for _, v := range ConfLst {
		_, err = toml.DecodeFile(v.Path, v.Obj)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func Reload(name string) (err error) {
	name = filepath.ToSlash(name)
	fmt.Printf("%+v\n", name)
	for _, v := range ConfLst {
		if find := strings.Contains(v.Path, name); find {
			_, err = toml.DecodeFile(v.Path, v.Obj)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}

func Listen(evtName string, callback func(string)) {

}
