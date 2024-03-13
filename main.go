package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

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
	link := "https://www.baidu.com/"
	log.Printf("Chrome visit page %s\n", link)

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(link),
		chromedp.Click(`document.querySelector("#s-top-loginbtn")`, chromedp.ByJSPath),
		chromedp.SendKeys(`#TANGRAM__PSP_11__userName`, "18737387054", chromedp.ByID),
		chromedp.SendKeys(`#TANGRAM__PSP_11__password`, "18737387054.c", chromedp.ByID),
		chromedp.Click(`document.querySelector("#TANGRAM__PSP_11__isAgree")`, chromedp.ByJSPath),
		chromedp.Click(`document.querySelector("#TANGRAM__PSP_11__submit")`, chromedp.ByJSPath),
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
	)
	log.Printf("err %s\n", err)
	log.Printf("htmlContent %s\n", htmlContent)
}
