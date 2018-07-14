package controllers

import(
	"github.com/astaxie/beego"
	"blog/conn"
	"blog/middleware"
	"time"
	// "strconv"
)

func (u *Users)DbLogin(phone string,password string)(token string,has bool){
	conn.Init()
	db:=conn.GetOrm()
	err:=db.Where("phone=?&&password=?",phone,password).First(u).Error
	if err!=nil{
		beego.Info(err)
		has=false
		return
	}
	has=true
	token,errt:=middleware.CreateJwtToken(u.Uid,u.Name,u.Phone)
	if errt!=nil{
		beego.Info(errt)
	}
	return
}
var limit=10
type UserController struct{
	beego.Controller
}
func GetLastUinx(day int64)(start int64,end int64){
	day=day-1
	timeStr := time.Now().Format("2006-01-02")
    t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	start=timeNumber-24*60*60*day
	end=timeNumber+24*60*60
    return 
}
func(c *UserController)GetFilterUserList(){
	users:=make([]Users,0)
	usergroup:=c.GetString("usergroup")
	created,_:=c.GetInt64("created")
	offset,_:=c.GetInt("page")
	beego.Info("created:",created)
	conn.Init()
	db:=conn.GetOrm()
	if usergroup=="" && created==0{
		db.Table("users").Offset(limit*(offset-1)).Limit(limit).Find(&users)
	}else if created!=0{
		start,end:=GetLastUinx(created)
		db.Table("users").Where("created BETWEEN ? AND ?",start,end).Offset(limit*(offset-1)).Limit(limit).Find(&users)
	}else if usergroup!=""{
		db.Table("users").Where("usergroup=?",usergroup).Offset(limit*(offset-1)).Limit(limit).Find(&users)
	}
	
	c.Data["json"]=users
	c.ServeJSON()
}
func (c *UserController)SearchUserName(){
	type uNames struct{
		Uid int64
		Name string
	}
	type cjson struct{
		ReJson
		Data []uNames
	}
	
	rjson:=&cjson{ReJson{0,"数据操作成功"},nil}
	name:=c.Ctx.Input.Query("name")
	beego.Info(name)
	conn.Init()
	db:=conn.GetOrm()
	us:=make([]uNames,0)
	db.Table("users").Select("uid,name").Where("name like ?","%"+name+"%").Find(&us)
	rjson.Data=us
	c.Data["json"]=rjson
	c.ServeJSON()
}