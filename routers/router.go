package routers

import (
	"train_system/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/index", &controllers.IndexController{})
    beego.Router("/create_imformation", &controllers.CreateImformationController{})
    beego.Router("/create_table", &controllers.CreateTableController{})
    beego.Router("/edit_imformation", &controllers.EditImformationController{})
    beego.Router("/edit_table", &controllers.EditTableController{})
    beego.Router("/generate_table", &controllers.GenerateTableController{})
}
