package models

import (
	"github.com/astaxie/beego"
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)
func DbTest(){
	mysqluser:=beego.AppConfig.String("mysqluser")
	mysqlpass:=beego.AppConfig.String("mysqlpass")
	mysqlurls:=beego.AppConfig.String("mysqlurls")
	mysqlport:=beego.AppConfig.String("mysqlport")
	mysqldb:=beego.AppConfig.String("mysqldb")
    mysqlStr:=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",mysqluser,mysqlpass,mysqlurls,mysqlport,mysqldb)
	db, err := gorm.Open("mysql", mysqlStr)
	db.LogMode(true)
	if err!=nil{
		panic(err)
	}
	// c:=&Category{}
	// c:=new(Category)
	// err=db.Table("category").First(c).Error
	// beego.Info(err,c)
	// db.CreateTable(&Users{})
	// u:=&Users{Created:time.Now().Unix()}
	// err=db.Create(u).Error
	// beego.Info(err)
	// us:=make([]Users,0)
	// db.Table("users").Find(&us)
	// beego.Info(us)
	// type hav struct{
	// 	Num int64
	// }
	// var h []hav
	// u3:=make([]Users,0)
	// err=db.Limit(10).Find(&u3).Select("sum(age) as num").Group("sex").Having("sum(age)>?",100).Scan(&h).Error
	// beego.Info(u3,h)

	// type ut struct{
	// 	uid int64
	// 	title string
	// 	name string
	// }
	// var us []ut
	// db.Table("users").Select("users.id as uid,users.name,topic.title").Joins("left join topic on topic.uid = users.id").Scan(&us)
	// beego.Info(us)
	// var ages []int64
	// users:=new(Users)
	// db.Find(&users).Pluck("age", &ages)
	// beego.Info(ages)
	
}

// type Users struct{
// 	Id int64
// 	Name string 
// 	Age int64
// 	Sex int64
// 	Created int64
// 	Phone int
// 	Password string
// }

// func (u *User)TableName()string{
// 	return "users"
// }

type Category struct{
	Id int64
	Title string
	Created time.Time `gorm:"index"`
	Views int64 `gorm:"index"`
	TopicTime time.Time `gorm:"index"`
	TopicCount int64
	TopicLastUser int64
}
type Topic struct{
	Id int64
	Uid int64
	Title string
	Content string `gorm:"size(5000)"`
	Attachment []string
	Created time.Time `gorm:"index"`
	Views int64 `gorm:"index"`
	Update time.Time `gorm:"index"`
	Auther string
	ReplayTime time.Time 
	ReplayCount int64
	ReplayLastUserId int64
}
func RegisterDB(){

}