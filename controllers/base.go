package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}
// 请求前必须的操作，判断用户登录
func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
	}
	fmt.Print("BaseController：Prepare：判断用户是否登录")
	this.Data["isLogin"] = this.isLogin
}
// 默认404页面
func (this *BaseController) Go404() {
	this.TplName = "404.tpl"
	return
}
// 默认500页面
func (this *BaseController) Go500() {
	this.TplName = "500.tpl"
	return
}
// 默认显示页面
func (this *BaseController) GoIndex() {
	this.TplName = "index.tpl"
	return
}


