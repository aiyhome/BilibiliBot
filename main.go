package main

import (
	"BilibiliBot/app"
	"BilibiliBot/event"
	"BilibiliBot/util/log"
)

func main() {
	log.Info("==================== 初始化 ====================")
	app.Run()
	listener := event.NewListener(func(evt *event.Event, data ...interface{}) {
		log.Info("ok", evt)
		log.Info(data)
	})
	app.Attach("test", listener)
	app.Emit("test")
	app.Detach("test", listener)
	app.Emit("test")
	log.Info("=================== 运行结束 ===================")
}
