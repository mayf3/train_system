package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["train_system/controllers/person:PersonController"] = append(beego.GlobalControllerRouter["train_system/controllers/person:PersonController"],
		beego.ControllerComments{
			"Show",
			`/show`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/person:PersonController"] = append(beego.GlobalControllerRouter["train_system/controllers/person:PersonController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/person:PersonController"] = append(beego.GlobalControllerRouter["train_system/controllers/person:PersonController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

}
