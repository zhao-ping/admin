package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"os"
	"io"
	"github.com/astaxie/beego/validation"
	// _ "github.com/astaxie/beego/session/mysql"
	// "strconv"
)

type HomeController struct {
	beego.Controller
}
func (c *HomeController)Admin(){
	c.TplName="index.html"
}
func(this *HomeController) First(){
	var ob Object
	//ajax普通模式上传数据
	err:=json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	beego.Info("获取的obj:",err,ob)
	// u:=Object{Name:"q",Age:23,Sex:"man"}
	this.Data["json"] = ob
	this.ServeJSON()
	
	
}
func(c *HomeController) Upload(){
	f, h, err := c.GetFile("file")
    if err != nil {
        beego.Info("getfile err ", err)
    }
    defer f.Close()
    c.SaveToFile("file", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	c.Data["string"]="上传成功！"
	c.ServeJSON();
}
func(c *HomeController) Upload2(){
	//formData上传图片
	files:=c.Ctx.Request.MultipartForm.File
	for _,file:=range files{
		for _,fsin := range file {
			fio,err:=fsin.Open()
			f,err:=os.OpenFile(fmt.Sprintf("static/upload/%s",fsin.Filename),os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			beego.Info(err)
			io.Copy(f, fio)
			f.Close()
		}
	}
	c.ServeJSON();
}
type Object struct {
	Name string
	Age int
	Sex string
}

type serverJson struct{
	Code int
	Message string
}

func(c *HomeController)Login(){
	// u:=new(Users)
	// f:=json.Unmarshal(c.Ctx.Input.RequestBody,&u)
	// beego.Info(u,f)
	// beego.Info(c.GetString("phone"))
	// phone,_:=strconv.ParseInt(c.GetInt("phone"),10,64)
	beego.Info(c.Ctx.Input.Query("phone"))
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
	if has {
		j=serverJson{0,"登录成功"}
		c.Ctx.SetCookie("token",token)
	}else{
		j=serverJson{0,"登录失败"}
		c.Ctx.SetCookie("token","",0)
	}
	c.Data["json"]=j
	c.ServeJSON()
}
func (t *HomeController) Get() {
	t.TplName="home/index.html"
	t.Data["Title"]="home"
	t.Data["Hello"]="我是一只小小鸟"
	t.Data["Istrue"]=false
	type User struct{
		Name string
		Age int
		Sex string
	}
	user:=&User{
		Name:"张三",
		Age:18,
		Sex:"man",
	}
	t.Data["User"]=user
	s:=[]int{1,2,3,4,5,6,7,8,9,0}
	t.Data["Array"]=s
	t.Data["Var"]=1
	t.Data["Html"]="<b>hello world!</b>"
	t.Data["Pipe"]="<b>hello world!</b>"
	t.Ctx.Output.Cookie("name","value")

	
}
func (t *HomeController) Home() {
	
}

