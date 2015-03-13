package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "train_system/routers"
)

const (
	kDataBaseUser     = "iladmin"
	kDataBasePassword = ""
	kDataBaseName     = "test"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", kDataBaseUser+":"+kDataBasePassword+"@/"+kDataBaseName+"?charset=utf8")
}

func main() {
	orm.RunCommand()
	beego.Run()
}
