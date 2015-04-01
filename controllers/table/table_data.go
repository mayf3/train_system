package table

import (
	"fmt"

	"train_system/utils"
)

// @router /create [get]
// @router /:table_id:int/setting [get]
func (this *TableController) Setting() {
	this.checkTable()
	_, err := this.getTableId()
	if err == nil{
		table := this.getPreTable()
		this.Data["form_pre_table"] = table
		this.Data["form_problem_number"] = table.ProblemNumber
	} else{
		this.Data["form_problem_number"] = 0
	}
	this.Data["base_problem_list"] = this.getProblemList()
	fmt.Println(this.Data["form_problem_number"])
	this.TplNames = "table/setting_table.tpl"
}

// @router /create [post]
// @router /:table_id:int/setting [post]
func (this *TableController) SettingPost() {
	this.checkTable()
	this.setting()
	this.Redirect("/", 302)
}

// @router /:table_id:int/edit [get]
func (this *TableController) Edit() {
	this.mustTable()
	table_id, _ := this.getTableId()
	this.Data["data[problem_list]"] = utils.GenerateProblemList(table_id)
	this.Data["data[total_status]"] = utils.GenerateTable(table_id)
	this.TplNames = "table/edit_table.tpl"
}

// @router /:table_id:int/show [get]
func (this *TableController) Show() {
	this.mustTable()
	table_id, _ := this.getTableId()
	this.Data["data[current_table]"], this.Data["data[hust_table"] = this.generate()
	this.Data["data[problem_list]"] = utils.GenerateProblemList(table_id)
	this.Data["data[total_status]"] = utils.GenerateTable(table_id)
	this.TplNames = "table/show_table.tpl"
}

// @router /:table_id:int/show [post]
func (this *TableController) ShowPost() {
	this.mustTable()
	table_id, _ := this.getTableId()
	this.submitHust()
	this.Redirect(fmt.Sprintf("/table/%d/show", table_id), 302)
}

// @router /:table_id:int/del [post]
func (this *TableController) DelPost() {
	this.mustTable()
	this.del()
	this.Redirect("/", 302)
}
