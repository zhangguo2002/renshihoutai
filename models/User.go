package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户信息
type Users struct {
	Id   int
	Name string `orm:"unique"` // 用户名唯一
	Pwd  string
}

//员工资料信息
type YuanGongZiLiao struct {
	Id         int `orm:"pk"`
	Name       string
	Department string
	Entrytime  string
	Salary     int
}

//部门信息
type BuMenXinXi struct {
	Id         int `orm:"pk"`
	Name       string
	Department string
	Level      string
}

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/renshihoutai?charset=utf8")

	orm.RegisterModel(new(Users), new(YuanGongZiLiao), new(BuMenXinXi)) // 注册模型，建立User类型对象，注册模型时，需要引入包

	orm.RunSyncdb("default", false, true)
}
