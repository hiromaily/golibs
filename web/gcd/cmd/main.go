package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
)

var debugger *gcd.Gcd
var path string
var dir string
var port string

func init() {
	switch runtime.GOOS {
	case "windows":
		flag.StringVar(&path, "chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "path to chrome")
		flag.StringVar(&dir, "dir", "C:\\temp\\", "user directory")
	case "darwin":
		flag.StringVar(&path, "chrome", "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "path to chrome")
		flag.StringVar(&dir, "dir", "/tmp/", "user directory")
	case "linux":
		flag.StringVar(&path, "chrome", "/usr/bin/chromium-browser", "path to chrome")
		flag.StringVar(&dir, "dir", "/tmp/", "user directory")
	}

	flag.StringVar(&port, "port", "9222", "Debugger port")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	flag.Parse()
	debugger = gcd.NewChromeDebugger()

	// add flag
	extensionPath := "--print-to-pdf=test12345.pdf"
	debugger.AddFlags([]string{extensionPath})

	// start process, specify a tmp profile path so we get a fresh profiled browser
	// set port 9222 as the debug port
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess() // exit when done

	target, err := debugger.NewTab()
	if err != nil {
		log.Fatalf("error opening new tab: %s\n", err)
	}

	//subscribe to page load
	target.Subscribe("Page.loadEventFired", func(targ *gcd.ChromeTarget, v []byte) {
		doc, err := target.DOM.GetDocument(-1, true)
		if err == nil {
			log.Printf("%s\n", doc.DocumentURL)
		}
		wg.Done() // page loaded, we can exit now
		// if you wanted to inspect the full response data, you could do that here
	})

	// get the Page API and enable it
	if _, err := target.Page.Enable(); err != nil {
		log.Fatalf("error getting page: %s\n", err)
	}

	// TODO:When using local file, `file:///Users/yasui.hiroki/Desktop/hy.png`
	navigateParams := &gcdapi.PageNavigateParams{Url: "http://www.veracode.com"}
	ret, _, _, err := target.Page.NavigateWithParams(navigateParams) // navigate
	if err != nil {
		log.Fatalf("Error navigating: %s\n", err)
	}

	log.Printf("ret: %#v\n", ret)
	wg.Wait() // wait for page load

	//TODO: print out as PDF
	generatePDF(target)

	//debugger.CloseTab(target)
}

func generatePDF(target *gcd.ChromeTarget) {
	dom := target.DOM
	page := target.Page
	doc, err := dom.GetDocument(-1, true)
	if err != nil {
		fmt.Printf("error getting doc: %s\n", err)
		return
	}

	debugger.ActivateTab(target)
	time.Sleep(1 * time.Second) // give it a sec to paint
	u, urlErr := url.Parse(doc.DocumentURL)
	if urlErr != nil {
		fmt.Printf("error parsing url: %s\n", urlErr)
		return
	}

	// TODO: PDF configuration from here
	fmt.Printf("Print out as pdf: %s\n", u.Host)
	printPDFParams := &gcdapi.PagePrintToPDFParams{
		Landscape:               false,
		DisplayHeaderFooter:     false,
		PrintBackground:         false,
		Scale:                   1,
		PaperWidth:              8.5,
		PaperHeight:             11,
		MarginTop:               0.4,
		MarginBottom:            0.4,
		MarginLeft:              0.4,
		MarginRight:             0.4,
		PageRanges:              "",
		IgnoreInvalidPageRanges: false,
		HeaderTemplate:          "",
		FooterTemplate:          "",
		PreferCSSPageSize:       false,
	}

	//data: Base64-encoded pdf data.
	data, err := page.PrintToPDFWithParams(printPDFParams)

	if err != nil {
		fmt.Printf("error PrintToPDFWithParams: %s\n", err)
		return
	}
	fmt.Println("data is:", data)

	//fmt.Printf("Taking screen shot of: %s\n", u.Host)
	//screenShotParams := &gcdapi.PageCaptureScreenshotParams{Format: "png", FromSurface: true}
	//img, errCap := page.CaptureScreenshotWithParams(screenShotParams)
	//if errCap != nil {
	//	fmt.Printf("error getting doc: %s\n", errCap)
	//	return
	//}
	//
	//imgBytes, errDecode := base64.StdEncoding.DecodeString(img)
	//if errDecode != nil {
	//	fmt.Printf("error decoding image: %s\n", errDecode)
	//	return
	//}
	//
	//f, errFile := os.Create(u.Host + ".png")
	//defer f.Close()
	//if errFile != nil {
	//	fmt.Printf("error creating image file: %s\n", errFile)
	//	return
	//}
	//f.Write(imgBytes)

	debugger.CloseTab(target)
}

func takeScreenShot(target *gcd.ChromeTarget) {
	dom := target.DOM
	page := target.Page
	doc, err := dom.GetDocument(-1, true)
	if err != nil {
		fmt.Printf("error getting doc: %s\n", err)
		return
	}

	debugger.ActivateTab(target)
	time.Sleep(1 * time.Second) // give it a sec to paint
	u, urlErr := url.Parse(doc.DocumentURL)
	if urlErr != nil {
		fmt.Printf("error parsing url: %s\n", urlErr)
		return
	}

	fmt.Printf("Taking screen shot of: %s\n", u.Host)
	screenShotParams := &gcdapi.PageCaptureScreenshotParams{Format: "png", FromSurface: true}
	img, errCap := page.CaptureScreenshotWithParams(screenShotParams)
	if errCap != nil {
		fmt.Printf("error getting doc: %s\n", errCap)
		return
	}

	imgBytes, errDecode := base64.StdEncoding.DecodeString(img)
	if errDecode != nil {
		fmt.Printf("error decoding image: %s\n", errDecode)
		return
	}

	f, errFile := os.Create(u.Host + ".png")
	defer f.Close()
	if errFile != nil {
		fmt.Printf("error creating image file: %s\n", errFile)
		return
	}
	f.Write(imgBytes)
	debugger.CloseTab(target)
}
