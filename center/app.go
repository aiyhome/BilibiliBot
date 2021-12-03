package app

import (
    "BilibiliBot/conf"
    "BilibiliBot/util/log"
    "fmt"
    "runtime"
)

type App struct {
}

func (app *App) Run() {
    app.initMaxProcs()
    app.initConfig()
    app.initLogger()
}

func (app *App) initMaxProcs() {
    runtime.GOMAXPROCS(maxProcs)
}

func (app *Ego) initConfig() {
    conf.Load()

    fmt.Printf("%+v\n", cfg)
    fmt.Printf("%+v\n", conf.Platforms)
    fmt.Printf("%+v\n", conf.Users)
}

func (app *Ego) initLogger() {
    cfg := conf.Common.Log
    log.Init(cfg.Directory, cfg.FilePattern, cfg.FileLink, cfg.ConsoleOnly)
}
