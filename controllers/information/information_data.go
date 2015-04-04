package information

import (
	"fmt"
)

// @router /:information_id:int/edit [get]
func (this *InformationController) Edit() {
	this.checkTable()
	this.Data["other_all_person"] = this.getAllPerson()
	if this.checkInformation() == true {
		this.Data["form_pre_information"] = this.getInformation()
		this.Data["form_pre_member"] = this.getMember()
	} else {
		this.Data["form_pre_member"] = [3]int{0, 0, 0}
	}
	this.Data["form_problem_status"] = this.getProblemStatus()
	fmt.Println(this.Data["form_problem_status"])
	table_id, _ := this.getTableId()
	this.Data["url_current_table"] = fmt.Sprintf("/table/%d", table_id)
	this.TplNames = "information/edit_information.tpl"
}

// @router /:information_id:int/edit [post]
func (this *InformationController) EditPost() {
	this.checkTable()
	table_id, _ := this.getTableId()
	this.submitInformation()
	this.Redirect(fmt.Sprintf("/table/%d/edit", table_id), 302)
}
