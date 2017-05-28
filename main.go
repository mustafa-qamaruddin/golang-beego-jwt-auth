package main

import (
	_ "passapp-engine-api/routers"
	"github.com/astaxie/beego"
	"strconv"
    "os"
	"fmt"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {
		port, err := strconv.Atoi(os.Getenv("PORT"))
        if err == nil {
				fmt.Println(port)
                beego.BConfig.Listen.HTTPPort = port
        }
	}
	beego.Run()
}
