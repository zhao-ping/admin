package controllers
import(
)
type Article struct{
	Aid int64 `gorm:"primary_key"`
	Uid int64
	Title string
	Content string
	Attachment string
	Created int64
	State int64
	Views int64
	Update int64
	Auther string
	ReplayCount int64 //`grom:"column:replaycount"`
	ReplayLastUserId int64 
	RecommendText string 
}
type ReJson struct{
	Code int64
	Message string
}
type Users struct{
	Uid int64
	Name string `form:"name"`
	Avatar string
	Age int64 
	Sex int64
	Created int64
	Phone string `form:"phone"`
	Password string `form:"password"`
	Usergroup int64
}