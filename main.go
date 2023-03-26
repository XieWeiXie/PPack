package main

import "github.com/XieWeiXie/PPack/Pack"

var (
	APPName  string
	ICONName string
)

func main() {
	app := Pack.NewMacApplication(APPName, ICONName)
	_ = Pack.Do(app)
}
