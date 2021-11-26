package main

import (
	conf "BilibiliBot/conf"
	// spider "BilibiliBot/spider"

	log "BilibiliBot/util/log"
	"fmt"
	_ "github.com/dchest/captcha"
	_ "github.com/hashicorp/go-plugin"
	_ "io"
	_ "os"
	// "time"
)

func init() {
	log.I("==================== 初始化 ====================")
}

func main() {
	log.I("=============== 启动 spider 测试 ===============")

	conf.Load()

	// spider.SeleniumTest()
	fmt.Printf("%+v\n", conf.Common)
	log.I("=================== 运行结束 ===================")
}
