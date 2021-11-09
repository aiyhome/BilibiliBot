package spider

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func SeleniumTest() {

	const (
		seleniumPath = `./chromedriver`
		port         = 9515
	)

	opts := []selenium.ServiceOption{}

	selenium.SetDebug(true)

	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		return
	}

	// ////server关闭之后，chrome窗口也会关闭
	// defer service.Stop()

	//链接本地的浏览器 chrome
	caps := selenium.Capabilities{
		//"browserName": "/Applications/Google Chrome Dev.app/Contents/MacOS/Google Chrome Dev",
		"browserName": "Google Chrome Dev",
	}

	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		// Path:  "/Applications/Google Chrome Dev.app/Contents/MacOS/Google Chrome Dev",
		Path: "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe",
		Args: []string{
			//静默执行请求
			"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	url := "https://passport.bilibili.com/login"
	w_b1, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}

	err = w_b1.Get(url)
	//获取网页源码
	pageSource, err := w_b1.PageSource()
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
	fmt.Print(pageSource)

	defer service.Stop() // 停止chromedriver
	defer w_b1.Quit()    // 关闭浏览器

	return
}
