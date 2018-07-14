package controllers

import(
	"github.com/astaxie/beego"
	"blog/conn"
	"encoding/json"
	"time"
	"blog/middleware"
	"strconv"
	"fmt"
)


type ArticleController struct{
	beego.Controller
}
func (c *Article)AfterCreate() (err error){
	beego.Info("文章ID(Aid)",c.Aid)
	return
}
func (c *ArticleController)Get(){
	c.TplName="home/addArticle.html"
}

func (c *ArticleController)AddArticle(){
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.106:8080")
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Credentials", "true")
	type rejson struct{
		ReJson
		Aid int64
	}
	rjson:=new(rejson)
	rjson.Code=0
	rjson.Message="添加文章成功"
	article:=new(Article)
	// aid:=c.GetString("aid")
	err:=json.Unmarshal(c.Ctx.Input.RequestBody,&article)
	if err!=nil{
		beego.Info("json unmarshal errr:",err)
	}
	
	beego.Info("article:",article)
	if(article.Uid==0){
		token:=c.Ctx.GetCookie("token")
		u,tokenErr:=middleware.ParseJwtToken(token)
		if tokenErr!=nil{
			beego.Error("tokenErr:",tokenErr)
			rjson.Message="解析token失败"
		}
		article.Uid=u.Id
		article.Auther=u.Name
	}
	
	
	if article.Created==0{
		article.Created=time.Now().Unix()
		article.Update=time.Now().Unix()
	}else{
		article.Update=article.Created
	}
	conn.Init()
	db:=conn.GetOrm()
	dberr:=db.Table("article").Create(&article)
	if dberr!=nil{
		beego.Error(dberr)
	}
	rjson.Aid=article.Aid
	c.Data["json"]=rjson
	c.ServeJSON()
}
func (c *ArticleController) ListPage(){
	c.TplName="home/list.html"
}
type aReJson struct{
	ReJson
	Data []Article
	TotalCount int
}
func (c *ArticleController) GetSingleArticle(){
	aid:=c.Ctx.Input.Query("aid")
	beego.Info("aid:",aid)
	type rArticle struct{
		Article
		Avatar string
	}
	type singleArticle struct{
		ReJson
		Data rArticle
	}
	rjson:=new(singleArticle)
	conn.Init()
	db:=conn.GetOrm()
	db.Table("article").Select("*").Where("aid=?",aid).Joins("left join users on users.uid = article.uid").First(&rjson.Data)
	c.Data["json"]=rjson
	c.ServeJSON()
}
func (c *ArticleController) GetArticleList(){
	auther:=c.GetString("auther")
	state:=c.GetString("state")
	filter:=""
	s:=""
	if auther!=""{
		filter=fmt.Sprintf("auther='%s'",auther) 
		s=" && "
	}
	if state!="" {
		filter=fmt.Sprint(filter,fmt.Sprintf(" %s state=%s",s,state)) 
	}
	
	limit,_:=c.GetInt("limit",10)
	if limit==0{
		limit=10
	}
	page,_:=strconv.Atoi(c.GetString("page")) 
	 var totalCount int
	trjson:=new(aReJson)
	trjson.Data,totalCount=ArticleList(page,limit,filter,"article")
	beego.Info("totalCount:",totalCount)
	if (totalCount-limit*page)<=0{
		trjson.Code=1
		trjson.Message="数据已经加载完毕"
	}
	trjson.TotalCount=totalCount
	c.Data["json"]=trjson
	c.ServeJSON();
}
func ArticleList(page int,limit int,filter string,tbName string) (lists []Article,totalCount int){
	conn.Init()
	db:=conn.GetOrm()
	beego.Info("fun offset:",page)
	beego.Info("fun limit:",limit)
	errdb:=db.Table(tbName).Where(filter).Count(&totalCount).Offset((page-1)*limit).Limit(limit).Find(&lists).Error
	if errdb!=nil{
		beego.Info(errdb)
	}
 	return
}
func (c *ArticleController)EditArticle(){
	rjson:=ReJson{Code:0,Message:"操作成功"}
	content:=c.Ctx.Input.Query("Content")
	title:=c.Ctx.Input.Query("Title")
	aid:=c.Ctx.Input.Query("Aid")
	conn.Init()
	db:=conn.GetOrm()
	err:=db.Table("article").Where("aid=?",aid).Updates(&Article{Content:content,Title:title}).Error
	if err!=nil{
		beego.Error("eidt article db err:",err)
		rjson.Message="数据库操作出错"
		rjson.Code=1
	}
	c.Data["json"]=rjson
	c.ServeJSON()
}
func (c*ArticleController)ChangeArticleState(){
	rjson:=ReJson{Code:0,Message:"修改状态成功"}
	recommendText:=c.Ctx.Input.Query("RecommendText")
	state,_:=c.GetInt64("State")
	aid:=c.Ctx.Input.Query("Aid")
	beego.Info("ChangeArticleState:",aid,state,recommendText)
	conn.Init()
	db:=conn.GetOrm()
	err:=db.Table("article").Where("aid=?",aid).Updates(map[string]interface{}{"recommend_text":recommendText,"state":state}).Error
	if err!=nil{
		beego.Error("ChangeArticleState db err:",err)
		rjson.Code=1
		rjson.Message="数据库操作错误"
	}
	c.Data["json"]=rjson
	c.ServeJSON()
}