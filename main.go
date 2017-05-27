package main

import (
	_ "passapp-engine-api/routers"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
