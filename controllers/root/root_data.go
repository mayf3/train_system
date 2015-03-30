package root

// @router / [get]
func (this *RootController) Index() {
	this.Data["url[create_table]"] = "/table/create"
	this.Data["show[all_table]"] = this.getAllTable()
	this.TplNames = "root/index.tpl"
}
