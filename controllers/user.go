package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

}

func (c *UserController) Login() {
	phone:=c.Ctx.Input.Param(":phone")
	password:=c.Ctx.Input.Param(":password")
	c.Ctx.WriteString("账号：" + phone)
	c.Ctx.WriteString("密码：" + password)
}
