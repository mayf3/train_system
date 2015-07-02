package routers

import (
	"github.com/astaxie/beego"

	. "train_system/controllers/root"
)

func init() {
	beego.Include(&RootController{})
	//
	//	person := beego.NewNamespace("/person", beego.NSInclude(&PersonController{}))
	//	beego.AddNamespace(person)
	//
	//	table := beego.NewNamespace("/table",
	//		beego.NSInclude(&TableController{}),
	//		beego.NSNamespace("/:table_id:int/information",
	//			beego.NSInclude(&InformationController{})))
	//	beego.AddNamespace(table)
	//
	//	user := beego.NewNamespace("/user", beego.NSInclude(&UserController{}))
	//	beego.AddNamespace(user)
}
