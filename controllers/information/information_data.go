package information

import (
	"fmt"
)

// @router /create [get]
// @router /:information_id:int/edit [get]
func (this *InformationController) Edit() {
	this.checkTable()
	this.Data["form[all_person]"] = this.getAllPerson()
	if this.checkInformation() {
		this.Data["form[pre_inforamtion]"] = this.getInformation()
		this.Data["form[member]"] = this.getMember()
		this.Data["form[problem_status]"] = this.getProblemStatus()
	} else {
		this.Data["form[member]"] = [3]int{0, 0, 0}
	}
	this.TplNames = "information/edit_information.tpl"
}

// @router /create [post]
// @router /:information_id:int/edit [post]
func (this *InformationController) EditPost() {
	this.checkTable()
	table_id, _ := this.getTableId()
	this.submitInformation()
	this.Redirect(fmt.Sprintf("/table/%d/edit", table_id), 302)
}
