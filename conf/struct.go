package conf

/*
Convert TOML to Go struct:
source:
https://github.com/xuri/toml-to-go

online tool:
https://xuri.me/toml-to-go/
*/
type PlatformsConf struct {
	WebHeaders struct {
		Accept         string `toml:"Accept"`
		UserAgent      string `toml:"UserAgent"`
		AcceptLanguage string `toml:"AcceptLanguage"`
		AcceptEncoding string `toml:"AcceptEncoding"`
		Cookie         string `toml:"Cookie"`
	} `toml:"WebHeaders"`
	AppHeaders struct {
		Accept         string `toml:"Accept"`
		UserAgent      string `toml:"UserAgent"`
		AcceptLanguage string `toml:"AcceptLanguage"`
		AcceptEncoding string `toml:"AcceptEncoding"`
		Buvid          string `toml:"Buvid"`
		DisplayID      string `toml:"DisplayId"`
		DeviceGUID     string `toml:"DeviceGuid"`
		DeviceID       string `toml:"DeviceId"`
		Connection     string `toml:"Connection"`
		Cookie         string `toml:"Cookie"`
	} `toml:"AppHeaders"`
	AppParams struct {
		Actionkey string `toml:"Actionkey"`
		Appkey    string `toml:"Appkey"`
		Build     string `toml:"Build"`
		Device    string `toml:"Device"`
		MobiApp   string `toml:"MobiApp"`
		Platform  string `toml:"Platform"`
		AppSecret string `toml:"AppSecret"`
	} `toml:"AppParams"`
	TvHeaders struct {
		Accept         string `toml:"Accept"`
		UserAgent      string `toml:"UserAgent"`
		AcceptLanguage string `toml:"AcceptLanguage"`
		AcceptEncoding string `toml:"AcceptEncoding"`
		Cookie         string `toml:"Cookie"`
	} `toml:"TvHeaders"`
	TvParams struct {
		Actionkey string `toml:"Actionkey"`
		Appkey    string `toml:"Appkey"`
		Build     string `toml:"Build"`
		Device    string `toml:"Device"`
		MobiApp   string `toml:"MobiApp"`
		Platform  string `toml:"Platform"`
		AppSecret string `toml:"AppSecret"`
	} `toml:"TvParams"`
}

type UsersConf struct {
	Users []struct {
		UserName     string `toml:"UserName"`
		Password     string `toml:"Password"`
		AccessKey    string `toml:"AccessKey"`
		Cookie       string `toml:"Cookie"`
		Csrf         string `toml:"Csrf"`
		UID          string `toml:"Uid"`
		RefreshToken string `toml:"RefreshToken"`
	} `toml:"Users"`
}

type CommonConf struct {
	Log struct {
		ConsoleOnly bool   `toml:"ConsoleOnly"`
		Directory   string `toml:"Directory"`
		FilePattern string `toml:"FilePattern"`
		FileLink    string `toml:"FileLink"`
	} `toml:"Log"`
}
