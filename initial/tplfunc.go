package initial

import (
	"github.com/lock-upme/beegoblog/utils"
	"html/template"
	"net/http"
	"github.com/astaxie/beego"
)

func InitTplFunc() {
	beego.AddFuncMap("date_mh", utils.GetDateMH)
	beego.AddFuncMap("date", utils.GetDate)
	beego.AddFuncMap("avatar", utils.GetGravatar)

	beego.ErrorHandler("404", Page_not_found)
	beego.ErrorHandler("500", Error_not_handle)
}

func Page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
func Error_not_handle(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("500.tpl").ParseFiles("views/500.tpl")
	data := make(map[string]interface{})
	data["content"] = "error not handle"
	t.Execute(rw, data)
}