package table

import (
	"fmt"

	"github.com/mayf3/train_system/utils"
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
	this.TplName = "table/setting_table.tpl"
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
	this.Data["show_total_status"] = utils.GenerateTable(table_id)
	this.Data["other_problem_list"] = utils.GenerateProblemList(table_id)
	this.Data["url_current_table"] = fmt.Sprintf("/table/%d", table_id)
	this.Data["url_create_information"] = fmt.Sprintf("/table/%d/information/0/edit", table_id)
	this.TplName = "table/edit_table.tpl"
}

// @router /:table_id:int/show [get]
func (this *TableController) Show() {
	this.mustTable()
	table_id, _ := this.getTableId()
	this.Data["show_current_table"] = this.getTable()
	this.Data["show_hust_table"] = this.generate()
	this.Data["show_total_status"] = utils.GenerateTable(table_id)
	this.Data["other_problem_list"] = utils.GenerateProblemList(table_id)
	this.TplName = "table/show_table.tpl"
}

// @router /:table_id:int/show [post]
func (this *TableController) ShowPost() {
	this.mustTable()
	table_id, _ := this.getTableId()
	this.submitHust()
	this.Redirect(fmt.Sprintf("/table/%d/show", table_id), 302)
}

// @router /:table_id:int/del [get]
func (this *TableController) DelPost() {
	this.mustTable()
	this.del()
	this.Redirect("/", 302)
}
