package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["train_system/controllers/root:RootController"] = append(beego.GlobalControllerRouter["train_system/controllers/root:RootController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

}
