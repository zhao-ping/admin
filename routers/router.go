package routers

import (
    "blog/controllers"
    "github.com/astaxie/beego/context"
    "github.com/astaxie/beego"
    "blog/middleware"
// "net/http"
// "html/template"
)
// func page_not_found(rw http.ResponseWriter, r *http.Request){
//         t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/404.html")
//         data :=make(map[string]interface{})
//         data["content"] = "page not found"
//         t.Execute(rw, data)
//     }
func init() {
    //页面错误
    beego.ErrorController(&controllers.ErrorController{})
    beego.Router("/", &controllers.HomeController{},"get:Admin")

    //正常路由
    beego.Router("/home", &controllers.HomeController{})
    beego.Router("/first", &controllers.HomeController{},"post:First")
    beego.Router("/upload", &controllers.HomeController{},"post:Upload")
    beego.Router("/upload2", &controllers.HomeController{},"post:Upload2")
    // 登录
    beego.Router("/login", &controllers.LoginController{},"get:Get;post:Login")
    beego.Router("/register",&controllers.LoginController{},"get:RegisterPage;post:Register")

    //users
    beego.Router("/user/GetFilterUserList",&controllers.UserController{},"get:GetFilterUserList")
    beego.Router("/searchUserName",&controllers.UserController{},"get:SearchUserName")

    // article
    beego.Router("/addArticle",&controllers.ArticleController{},"get:Get;post:AddArticle")
    beego.Router("/list",&controllers.ArticleController{},"get:ListPage")
    beego.Router("/getArticleList",&controllers.ArticleController{},"get:GetArticleList")
    beego.Router("/getSingleArticle",&controllers.ArticleController{},"get:GetSingleArticle")
    beego.Router("/editArticle",&controllers.ArticleController{},"get:EditArticle")
    beego.Router("/changeArticleState",&controllers.ArticleController{},"get:ChangeArticleState")

    //web 列表分页插件
    beego.Router("/web/list/?:page([0-9]+)",&controllers.WebListController{},"get:Get;post:GetList")
 //路由过滤
    var FilterUser = func(ctx *context.Context) {
    token :=  ctx.Input.Cookie("token")
    // if len(token)==0{
    //     resp := response(2, -1, errms.ErrorLoseToken, nil)
    //     ctx.Output.JSON(&resp, true, true)
    //     return
    // }
    jwtClaims,err:=middleware.ParseJwtToken(token)
    if err!=nil{
        beego.Info(err)
    }
    beego.Info("解密后解构：",jwtClaims)
    return
}
	beego.InsertFilter("/user",beego.BeforeRouter,FilterUser)
    
    beego.InsertFilter("*",beego.BeforeRouter,func(ctx *context.Context) {
        ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	    ctx.ResponseWriter.Header().Add("Access-Control-Allow-Credentials", "true")
    })
    // beego.InsertFilter("*", beego.BeforeRouter, Allow(&Options{
    //     AllowAllOrigins:  true,
    //     AllowMethods:     []string{"*"},
    //     AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
    //     ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
    //     AllowCredentials: true,
    // }))
   
	
}
