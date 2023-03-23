package main

import "github.com/XieWeiXie/PPack/Pack"

func main() {
	app := Pack.NewMacApplication("douyin", "douyin_512.png")
	Pack.Do(app)
}
