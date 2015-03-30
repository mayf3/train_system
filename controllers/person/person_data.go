package person

//@router / [get]
func (this *PersonController) Show() {
	this.Data["data[all_person]"] = this.getAllPerson()
	this.TplNames = "person/person.tpl"
}

// @router /add [post]
func (this *PersonController) Add() {
	this.add()
	this.Redirect("/person", 302)
}

// @router /edit [post]
func (this *PersonController) Edit() {
	this.edit()
	this.Redirect("/person", 302)
}
