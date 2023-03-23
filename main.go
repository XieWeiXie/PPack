package main

import "github.com/XieWeiXie/PPack/Pack"

func main() {
	app := Pack.NewMacApplication("youtube", "douyin_512.png")
	Pack.Do(app)
}
