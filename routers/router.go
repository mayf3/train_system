package routers

import (
	"github.com/astaxie/beego"

	. "train_system/controllers/information"
	. "train_system/controllers/person"
	. "train_system/controllers/root"
	. "train_system/controllers/table"
)

func init() {
	beego.Include(&RootController{})

	person := beego.NewNamespace("/person", beego.NSInclude(&PersonController{}))
	beego.AddNamespace(person)

	table := beego.NewNamespace("/table",
		beego.NSInclude(&TableController{}),
		beego.NSNamespace("/:table_id:int",
			beego.NSInclude(&InformationController{})))

	beego.AddNamespace(table)
}
