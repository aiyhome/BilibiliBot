package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	confFilepath  string
	usersFilepath string

	Common = &CommonConf{}
	Users  = &UsersConf{}
)

func init() {
	flag.StringVar(&confFilepath, "conf", "./conf/toml/conf.toml", "-conf path")
	flag.StringVar(&usersFilepath, "users", "./conf/toml/users.toml", "-users path")
}

// Init init conf
func Load() (err error) {
	_, err = toml.DecodeFile(confFilepath, &Common)
	if err != nil {
		log.Fatal(err)
	}
	_, err = toml.DecodeFile(usersFilepath, &Users)
	if err != nil {
		log.Fatal(err)
	}
	return
}
