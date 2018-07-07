package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

/*
1、添加一个关联user的order
*/
func AddOrder() {
	o := orm.NewOrm()
	order := Order{Name: "西瓜订单"}
	order.User = &(User{Id: 1})
	id, e := o.Insert(order)
	if e != nil {
		beego.Info("添加order错误：", e)
	}

	beego.Info("添加order成功：", id)
}

/*
1、查询多条order数据
*/
func findOrders() {
	var orders [] Order
	o := orm.NewOrm()
	qs := o.QueryTable("Order") //指定查询哪个表
	/*user__id两个下划线，代表，user.id*/
	n, e := qs.Filter("user__id", 1).All(orders)//指定查询user.id为1的所有订单，查询所有，把查询的数据放在orders内
	if e != nil {
		beego.Info("v错误：", e)
	}

	beego.Info("查询多条order数据：n:", n)
}
