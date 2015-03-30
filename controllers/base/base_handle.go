package base

func (this *BaseController) prepare() {
	var (
		tmp string
	)
	this.Data["Permission"] = "guest"
	tmp = this.Input().Get("admin")
	if tmp == "true" {
		this.Data["Permission"] = "admin"
	}
}
