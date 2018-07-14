package conn
import(
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db*gorm.DB
func Init(){
	mysqluser:=beego.AppConfig.String("mysqluser")
	mysqlpass:=beego.AppConfig.String("mysqlpass")
	mysqlurls:=beego.AppConfig.String("mysqlurls")
	mysqlport:=beego.AppConfig.String("mysqlport")
	mysqldb:=beego.AppConfig.String("mysqldb")
	mysqlStr:=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",mysqluser,mysqlpass,mysqlurls,mysqlport,mysqldb)
	var err error
	db, err = gorm.Open("mysql", mysqlStr)
	db.LogMode(true)
	if err!=nil{
		panic(err)
	}
}
// 获取链接
func GetOrm() *gorm.DB {
	return db
}
