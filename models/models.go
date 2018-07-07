package models

import "github.com/astaxie/beego/orm"
import _ "github.com/go-sql-driver/mysql"

/*
1、rel(fk)声明一对多
*/
type User struct {
	Id     int
	Name   string    `orm:size(100)`
	Orders *[] Order `rom:rel(fk)`
}

/*
1、reverse(many)设置多对一
*/
type Order struct {
	Id   int
	Name string `orm:size(100)`

	User *User `orm:reverse(many)`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "user:111111@/test03?charset=utf8",30)
	orm.RegisterModel(new(User) )

	orm.RunSyncdb("default", false, true)

}
