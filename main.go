package main

import (
	_ "Beego_Web/routers"
	"github.com/astaxie/beego"
	_ "Beego_Web/models"
)



func main() {

	/*
	1、静态路径的设置
	2、可以通过localhost:8080/download/demo.apk访问apk文件下的其他文件
	*/
	beego.SetStaticPath("download","apk")

	beego.Run()
	//models.AddUser()
}
