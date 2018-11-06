package main

import (
	_ "enzoism/goblog/routers"
	_ "enzoism/goblog/initial"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

