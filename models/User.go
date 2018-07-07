package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*添加用户到数据库*/
func AddUser() {
	orm := orm.NewOrm()
	user := User{Name: "chenliang"}
	id, error := orm.Insert(user)
	if error != nil {
		beego.Info("添加用户错误：", error)
	}

	beego.Info("添加用户成功：", id)

}

/*
从数据库查询
*/
func FindUser() {
	orm := orm.NewOrm()
	user := User{Id: 1}
	error := orm.Read(&user)
	if error != nil {
		beego.Info("查询用户错误：", error)
	}

	beego.Info("查询用户成功：", user)
}

/*
更新
*/
func UpdateUser() {
	o := orm.NewOrm()
	user := User{Id: 1}
	_, error := o.Update(&user)
	if error != nil {
		beego.Info("更新用户错误：", error)
	}

	beego.Info("更新用户成功：", user)
}

/*
删除
*/
func DeleteUser() {
	o := orm.NewOrm()
	user := User{Id: 1}
	_, error := o.Delete(&user)
	if error != nil {
		beego.Info("更新用户错误：", error)
	}

	beego.Info("更新用户成功：", user)
}
