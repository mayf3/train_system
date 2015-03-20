package routers

import (
	"github.com/astaxie/beego"
	"train_system/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/create_information", &controllers.CreateImformationController{})
	beego.Router("/create_table", &controllers.CreateTableController{})
	beego.Router("/edit_information", &controllers.EditImformationController{})
	beego.Router("/edit_table", &controllers.EditTableController{})
	beego.Router("/generate_table", &controllers.GenerateTableController{})
	beego.Router("/action", &controllers.ActionController{})
}
