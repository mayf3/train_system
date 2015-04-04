package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"Setting",
			`/create`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"Setting",
			`/:table_id:int/setting`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"SettingPost",
			`/create`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"SettingPost",
			`/:table_id:int/setting`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"Edit",
			`/:table_id:int/edit`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"Show",
			`/:table_id:int/show`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"ShowPost",
			`/:table_id:int/show`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["train_system/controllers/table:TableController"] = append(beego.GlobalControllerRouter["train_system/controllers/table:TableController"],
		beego.ControllerComments{
			"DelPost",
			`/:table_id:int/del`,
			[]string{"post"},
			nil})

}
