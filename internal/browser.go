package internal

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"time"
)

func InitBrowser() *rod.Browser {
	log.Printf("Initializing browser...")
	// Even you forget to close, rod will close it after main process ends.
	l := launcher.New()
	//l.Headless(false)
	// For more info: https://pkg.go.dev/github.com/go-rod/rod/lib/launcher
	u := l.MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	browser.Timeout(10 * time.Second)
	return browser
}
