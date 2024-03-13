package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

// 模拟百度登陆 跳不过极验
// sdfasfas123 zxcasdqwe123

func main() {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=true"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	chromeContext, _ := chromedp.NewExecAllocator(context.Background(), options...)
	chromeCtx, cancel := chromedp.NewContext(chromeContext, chromedp.WithLogf(log.Printf))

	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 20*time.Minute)
	defer cancel()
	link := "https://www.o7c25b.com:8553/"
	log.Printf("Chrome visit page %s\n", link)

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(link),
		chromedp.SendKeys(`document.querySelector("#popoverCon > div:nth-child(1) > div > div > div > input")`, "sdfasfas123", chromedp.ByJSPath),
		chromedp.SendKeys(`document.querySelector("#popoverCon > div:nth-child(1) > div > div > div > div.login_pwdBox__7FBSc > input")`, "zxcasdqwe123", chromedp.ByJSPath),
		chromedp.Click(`document.querySelector("#popoverCon > div:nth-child(1) > div > div > div > div.image_imageContainer__i7Z8e.image_cover__hOzAW.login_loginBtn__lEAIQ > div")`, chromedp.ByJSPath),
	)
	log.Printf("err %s\n", err)
	log.Printf("htmlContent %s\n", htmlContent)
}
