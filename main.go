package main

import (
    "github.com/astaxie/beego"
    // "github.com/astaxie/beego/logs"
    _"blog/routers"

)



func main() {
    // beego.SetStaticPath("/static","static")
    // models.Init()
    beego.SetLogger("file", `{"filename":"log/test.log"}`)
    beego.Run()
}

