package main

import (
	"github.com/webview/webview"
)

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
	w.Run()
}
