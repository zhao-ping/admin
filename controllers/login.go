package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"blog/conn"
	// "encoding/json"
	"time"
)
type LoginController struct{
	beego.Controller
}
func(l * LoginController)Get(){
	l.TplName="login/login.html"
}
func(c * LoginController)Login(){
	phone:=c.GetString("phone")
	password:=c.GetString("password")
	valid:=validation.Validation{}
	valid.Phone(phone,"phone")
	valid.MaxSize(password,12,"password")
	valid.MinSize(password,6,"password")
	var j serverJson

	if valid.HasErrors(){
		for _, err := range valid.Errors {
			beego.Info(err.Key, err.Message)
			j.Message+=err.Key+":"+err.Message+";"
		}
		c.Data["json"]=j
		c.ServeJSON()
		return
	}
	token,has:=new(Users).DbLogin(phone,password)
	// c.SetSecureCookie("uid")
	maxAge:=1
	if has {
		j=serverJson{0,"登录成功"}
		// c.Ctx.SetCookie("token", token,maxAge*24*60*60,"/", "http://192.168.0.106:8080",false,false)
		c.Ctx.SetCookie("token",token,maxAge*24*60*60,"/")
	}else{
		j=serverJson{0,"登录失败"}
		c.Ctx.SetCookie("token","",-1,"/")
	}
	c.Data["json"]=j
	c.ServeJSON()
}

func (c *LoginController)RegisterPage(){
	c.TplName="login/register.html"
}
func (c *LoginController)Register(){
	rjson:=ReJson{Code:0,Message:"注册成功"}
	var count int
	var u Users
	if er:=c.ParseForm(&u);er!=nil{
		beego.Info(er)
	}
	u.Created=time.Now().Unix()
	conn.Init()
	db:=conn.GetOrm()
	errdb:=db.Table("users").Where("phone=?",u.Phone).Or("name=?",u.Password).Count(&count).Error
	beego.Info(u)
	if errdb!=nil{
		rjson.Code=1
		rjson.Message="数据库查询出错"
		beego.Info("errdb",errdb)
	}
	if(count==0){
		if createErr:=db.Table("users").Create(&u).Error;createErr!=nil{
			beego.Info(createErr)
			return
		}
		token,_:=new(Users).DbLogin(u.Phone,u.Password)
		c.Ctx.SetCookie("token",token,24*60*60,"/")

	}
	c.Data["json"]=rjson
	c.ServeJSON()
}