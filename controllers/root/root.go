package root

import (
	"github.com/astaxie/beego"
)

type RootController struct {
	beego.Controller
}

// @router / [get]
func (this *RootController) Index() {
	this.TplNames = "root/index.tpl"
}
