package controllers

import(
"github.com/astaxie/beego"
)

type ErrorController struct{
	beego.Controller
}

func (e* ErrorController)Error404(){
	e.Data["content"]="page not found!"
	e.TplName="err.html"
}
func (e* ErrorController)Error501(){
	e.Data["content"]="server error!"
	e.TplName="err.html"
}
func (e* ErrorController)ErrorDb(){
	e.Data["content"]="database is now down!"
	e.TplName="err.html"
}