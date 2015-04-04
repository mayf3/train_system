package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["train_system/controllers/information:InformationController"] = append(beego.GlobalControllerRouter["train_system/controllers/information:InformationController"],
		beego.ControllerComments{
			"Edit",
			`/create`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/information:InformationController"] = append(beego.GlobalControllerRouter["train_system/controllers/information:InformationController"],
		beego.ControllerComments{
			"Edit",
			`/:information_id([0-9]+)/edit`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/information:InformationController"] = append(beego.GlobalControllerRouter["train_system/controllers/information:InformationController"],
		beego.ControllerComments{
			"EditPost",
			`/create`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/information:InformationController"] = append(beego.GlobalControllerRouter["train_system/controllers/information:InformationController"],
		beego.ControllerComments{
			"EditPost",
			`/:information_id([0-9]+)/edit`,
			[]string{"post"},
			nil})

}
