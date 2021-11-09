package conf

import (
	"flag"
	// "fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	confFilepath  string
	usersFilepath string
	// Conf global
	Conf  = &Config{}
	Users = &UserList{}
)

// Config .
type Config struct {
	WebHeaders *web_headers
	AppHeaders *app_headers
	AppParams  *app_params
	TvHeaders  *tv_headers
	TvParams   *tv_params
}

type web_headers struct {
	Accept         string
	UserAgent      string
	AcceptLanguage string
	AcceptEncoding string
	Cookie         string
}

type app_headers struct {
	UserAgent      string
	AcceptEncoding string
	Buvid          string
	DisplayId      string
	DeviceGuid     string
	DeviceId       string
	AcceptLanguage string
	Accept         string
	Connection     string
	Cookie         string
}

type app_params struct {
	Actionkey string
	Appkey    string
	Build     string
	Device    string
	Mobi_app  string
	Platform  string
	AppSecret string
}

type tv_headers struct {
	Accept         string
	UserAgent      string
	AcceptLanguage string
	AcceptEncoding string
	Cookie         string
}

type tv_params struct {
	Actionkey string
	Appkey    string
	Build     string
	Device    string
	MobiApp   string
	Platform  string
	AppSecret string
}

type UserList struct {
	Users []user
}

type user struct {
	UserName     string
	Password     string
	AccessKey    string
	Cookie       string
	Csrf         string
	Uid          string
	RefreshToken string
}

func init() {
	flag.StringVar(&confFilepath, "conf", "./conf/conf.toml", "-conf path")
	flag.StringVar(&usersFilepath, "users", "./conf/users.toml", "-users path")
}

// Init init conf
func Init() (err error) {
	_, err = toml.DecodeFile(confFilepath, &Conf)
	if err != nil {
		log.Fatal(err)
	}
	_, err = toml.DecodeFile(usersFilepath, &Users)
	if err != nil {
		log.Fatal(err)
	}
	return
}
