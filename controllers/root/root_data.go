package root

// @router / [get]
func (this *RootController) Index() {
	this.Data["data[all_table]"] = this.getAllTable()
	this.TplNames = "root/index.tpl"
}
