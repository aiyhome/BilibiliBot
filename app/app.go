package app

import (
	"BilibiliBot/conf"
	"BilibiliBot/conf/file"
	. "BilibiliBot/event"
	"BilibiliBot/util/log"
	"fmt"
	"runtime"
)

var maxProcs = 2

var (
	dispatcher *Dispatcher
)

func initMaxProcs() {
	runtime.GOMAXPROCS(maxProcs)
}

func initConfig() {
	conf.Load()

	fmt.Printf("%+v\n", conf.Platforms)
	fmt.Printf("%+v\n", conf.Users)
}

func initLogger() {
	cfg := conf.Common.Log
	fmt.Printf("%+v\n", cfg)
	log.Init(cfg.Directory, cfg.FilePattern, cfg.FileLink, cfg.ConsoleOnly)
}

func init() {
	initMaxProcs()
	initConfig()
	initLogger()
	dispatcher = NewDispatcher()
}

func Run() {
	file.New("./toml/config.toml", true)
}

func Emit(name string, data ...interface{}) {
	if len(data) > 0 {
		dispatcher.Event(name, data[0])
	} else {
		dispatcher.Event(name, nil)
	}
}

func Attach(name string, listener Listener) {
	dispatcher.Attach(name, listener)
}

func Detach(name string, listener Listener) {
	dispatcher.Detach(name, listener)
}

func On(name string, listener Listener) {
	Attach(name, listener)
}

func Off(name string, listener Listener) {
	Detach(name, listener)
}
