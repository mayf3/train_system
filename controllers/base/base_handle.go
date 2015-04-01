package base

import(
	"fmt"
	"strings"
)

func (this *BaseController) Prepare() {
	this.Data["permission"] = "guest"
	fmt.Println(strings.HasPrefix(this.Ctx.Input.Uri(), "/admin"))
	if strings.HasPrefix(this.Ctx.Input.Uri(), "/admin") == true{
		this.Data["permission"] = "admin"
	}
}
