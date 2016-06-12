package root

// @router / [get]
func (this *RootController) Index() {
	this.Data["url_create_table"] = "/table/create"
	this.Data["show_all_table"] = this.getAllTable()
	this.TplName = "root/index.tpl"
}
