package user

import (
	"fmt"
)

// Register : user register
// @router /register [get]
func (this *UserController) Register() {

	if this.CheckLoginRedirect(false) {
		return
	}

	this.TplNames = "user/register.tpl"
}

// RegisterPost : register post form
// @router /register [post]
func (this *UserController) RegisterPost() {

	if this.CheckLoginRedirect(false) {
		return
	}

	this.register()
	this.Redirect("/user/enter", 302)
}

// Enter : user log in
// @router /enter [get]
func (this *UserController) Enter() {

	if this.CheckLoginRedirect(false) {
		return
	}

	this.TplNames = "user/enter.tpl"
}

// EnterPost : Enter post form
// @router /enter [post]
func (this *UserController) EnterPost() {

	if this.CheckLoginRedirect(false) {
		return
	}

	this.enter()
	this.Redirect("/", 302)
}

// LogOut: user log out
// @router /logout [get]
func (this *UserController) LogOut() {

	if this.CheckLoginRedirect(true) {
		return
	}

	this.logout()
	this.Redirect("/", 302)
}

// VerifyUser : verify user
// @router /verify/:code [get]
func (this *UserController) VerifyUser() {
	code := this.GetString(":code")
	fmt.Println(code)
	this.verifyUser(code)
	this.Redirect("/", 302)
}

// ResetPassword : reset password
// @router /reset/:code [get]
func (this *UserController) ResetPassword() {
	code := this.GetString(":code")
	this.Data["form_username"] = this.verifyResetPwd(code)
	this.TplNames = "user/reset_pwd.tpl"
}

// ResetPassword : reset password
// @router /reset/:code [post]
func (this *UserController) ResetPasswordPost() {
	code := this.GetString(":code")
	this.verifyResetPwdPost(code)
	this.Redirect("/user/enter", 302)
}
