package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"enzoism/goblog/models"
)

//列表
type ListArticleController struct {
	BaseController
}

func (this *ListArticleController) Get() {
	page, err1 := this.GetInt("p")
	title := this.GetString("title")
	keywords := this.GetString("keywords")
	status := this.GetString("status")
	if err1 != nil {
		page = 1
	}

	offset, err2 := beego.AppConfig.Int("pageoffset")
	if err2 != nil {
		offset = 9
	}

	condArr := make(map[string]string)
	condArr["title"] = title
	condArr["keywords"] = keywords
	if !this.isLogin {
		condArr["status"] = "1"
	} else {
		condArr["status"] = status
	}
	countArticle := models.CountArticle(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countArticle)
	_, _, art := models.ListArticle(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["art"] = art
	userLogin := this.GetSession("userLogin")
	if userLogin !="" {
		this.Data["isLogin"] = userLogin
		this.Data["isLogin"] = this.isLogin
	}
	this.TplName = "article.tpl"
}