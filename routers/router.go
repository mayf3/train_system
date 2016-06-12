package routers

import (
	"github.com/astaxie/beego"

	. "github.com/mayf3/train_system/controllers/information"
	. "github.com/mayf3/train_system/controllers/person"
	. "github.com/mayf3/train_system/controllers/root"
	. "github.com/mayf3/train_system/controllers/table"
)

func init() {
	beego.Include(&RootController{})

	person := beego.NewNamespace("/person", beego.NSInclude(&PersonController{}))
	beego.AddNamespace(person)

	table := beego.NewNamespace("/table",
		beego.NSInclude(&TableController{}),
		beego.NSNamespace("/:table_id:int/information",
			beego.NSInclude(&InformationController{})))
	beego.AddNamespace(table)

	admin := beego.NewNamespace("/admin",
			beego.NSInclude(&RootController{}),
			beego.NSNamespace("/person",
				beego.NSInclude(&PersonController{})),
			beego.NSNamespace("/table",
				beego.NSInclude(&TableController{}),
				beego.NSNamespace("/:table_id:int/information",
					beego.NSInclude(&InformationController{}))))
	beego.AddNamespace(admin)
}
