package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var site, res string
	url := "https://github.com/chromedp/chromedp"
	err = c.Run(ctxt, task(url, &site, &res))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("saved screenshot from search result listing `%s` (%s)", res, site)
}

func task(url string, site, res *string) chromedp.Tasks {
	var buf []byte
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#hovercard-aria-description`, chromedp.ByID),
		chromedp.Location(site),
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			return ioutil.WriteFile("screenshot.png", buf, 0644)
		}),
	}
}
