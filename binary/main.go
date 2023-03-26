package main

import (
	"embed"
	"github.com/webview/webview"
)

//go:embed bind.js
var jsFile embed.FS

var (
	URL string
)

const (
	defaultWidth = 1200
	defaultHigh  = 780
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetSize(defaultWidth, defaultHigh, webview.HintNone)
	w.Navigate(URL)
	f, err := jsFile.ReadFile("bind.js")
	if err != nil {
		panic("Should add bind.js")
	}
	w.Init(string(f))
	w.Run()
}
