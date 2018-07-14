package controllers
import(
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
	// "math"
)
type WebListController struct{
	beego.Controller
}
func (c *WebListController)Get(){
	offset,_:=strconv.Atoi(c.Ctx.Input.Param(":page"))
	limit:=5
	lists,totalCount:=ArticleList(offset,limit,"","article")
	PageStr:=GetPageStr(offset,limit,totalCount,"/web/list/")
	c.Data["PageStr"]=PageStr
	c.Data["Lists"]=lists
	c.TplName="web/list.html"
}
func (c *WebListController)GetList(){
	c.ServeJSON()
}

func GetPageStr(offset,limit,totalCount int,url string)(PageStr string){
	pageCount:=totalCount/limit
	if totalCount%limit>0{
		pageCount+=1
	}
	PageStr=fmt.Sprintf(`<div class="pageBox"><div class="page"><span>第%d页，共%d页</span>`,offset,pageCount)
	// PageStr=`<div class="pageBox"><div class="page"><span>第`+strconv.Itoa(offset)+`页，共`+strconv.Itoa(pageCount)+`页</span>`
	if offset>1{
		PageStr=fmt.Sprint(PageStr,fmt.Sprintf(`<a href="%s1">第一页</a><a href="%s%d">上一页</a>`,url,url,offset-1))
		// PageStr+=`<a href="`+url+`1">第一页</a>
		// <a href="`+url+strconv.Itoa(offset-1)+`">上一页</a>`
	}
    currentPageClass:=""
	for i:=offset-2;i<=offset+2;i++{
		if i>0 && i<=pageCount{
			if i==offset{
				currentPageClass="active"
			}else{
				currentPageClass=""
			}
			// PageStr+=`<a class="`+currentPageClass+`" href="`+url+strconv.Itoa(i)+`">`+strconv.Itoa(i)+`</a>`
			PageStr=fmt.Sprint(PageStr,fmt.Sprintf(`<a class="%s" href="%s%d">%d</a>`,currentPageClass,url,i,i))
		}
	}

	if offset<pageCount{
		PageStr=fmt.Sprint(PageStr,fmt.Sprintf(`<a href="%s%d">下一页</a><a href="%s%d">最后一页</a>`,url,offset+1,url,pageCount))
		// PageStr+=`<a href="`+url+strconv.Itoa(offset+1)+`">下一页</a>
		// <a href="`+url+strconv.Itoa(pageCount)+`">最后一页</a>`
	}
	PageStr=fmt.Sprint(PageStr,"</div></div>")
	// PageStr+=`</div></div>`
	return PageStr
}