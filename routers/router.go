package routers

import (
	"github.com/astaxie/beego"
	"train_system/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/person", &controllers.PersonController{})
	beego.Router("/person/:action", &controllers.PersonController{})
	beego.Router("/create_information", &controllers.EditInformationController{})
	beego.Router("/create_table", &controllers.CreateTableController{})
	beego.Router("/edit_information", &controllers.EditInformationController{})
	beego.Router("/edit_table", &controllers.EditTableController{})
	beego.Router("/generate_table", &controllers.GenerateTableController{})
	beego.Router("/action", &controllers.ActionController{})
}
