package routers

import (
	"enzoism/goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 配置默认的基础类
	beego.Router("/", &controllers.ListArticleController{})
	beego.Router("/404.html", &controllers.BaseController{}, "*:Go404")
	beego.Router("/500.html", &controllers.BaseController{}, "*:Go500")

	//
	beego.Router("/article", &controllers.ListArticleController{})
}
