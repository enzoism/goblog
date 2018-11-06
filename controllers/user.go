package controllers

import (
	"enzoism/goblog/models"

)
type LoginUserController struct {
	BaseController
}

func (this *LoginUserController) Get() {
	check := this.isLogin
	if check {
		this.Redirect("/index", 302)
	} else {
		this.TplName = "login.tpl"
	}
}

func (this *LoginUserController) Post() {
	phone := this.GetString("phone")
	password := this.GetString("password")

	if "" == phone {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写手机号"}
		this.ServeJSON()
	}

	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
	}

	err, user := models.LoginUser(phone, password)
	if err == nil {
		this.SetSession("userLogin", "1")
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功", "user": user}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}
type LogoutUserController struct {
	BaseController
}

func (this *LogoutUserController) Get() {
	this.DelSession("userLogin")
	this.Redirect("/article", 302)
}
