package base

import(
//	"fmt"
	"strings"
)

func (this *BaseController) Prepare() {
	this.Data["permission"] = "guest"
	if strings.HasPrefix(this.Ctx.Input.URI(), "/admin") == true{
		this.Data["permission"] = "admin"
	}
}
