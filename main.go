package main

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/mayf3/train_system/routers"
	_ "github.com/mayf3/train_system/initialize"
)

const (
	kDataBaseUser     = "iladmin"
	kDataBasePassword = ""
	kDataBaseName     = "test"
)

func init() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", kDataBaseUser+":"+kDataBasePassword+"@/"+kDataBaseName+"?charset=utf8&loc=Asia%2FShanghai")
}

func main() {
	orm.RunCommand()
	beego.Run()
}
