package main

import (
	conf "BilibiliBot/conf"
	// spider "BilibiliBot/spider"
	log "BilibiliBot/util/log"
	"fmt"
	_ "github.com/dchest/captcha"
	_ "io"
	_ "os"
)

func init() {
	log.I("==================== 初始化 ==================== ")
}

func main() {
	log.I("==================== 启动 spider 测试 ==================== ")
	// spider.SeleniumTest()
	conf.Init()
	fmt.Printf("%+v\n", conf.Users)
}
