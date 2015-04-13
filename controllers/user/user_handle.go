package user

import (
	"fmt"

	//	"github.com/astaxie/beego"

	"train_system/modules/user"
)

func (this *UserController) register() {
	email := this.GetString("user[email]")
	username := this.GetString("user[username]")
	password := this.GetString("user[password]")
	err := user.RegisterUser(username, email, password)
	if err != nil {
		this.Abort(fmt.Sprintf("%v", err))
	}
	code := user.CreateUserActiveCode(username, nil)
	fmt.Println(code)
}

func (this *UserController) enter() {
	var (
		username string
		password string
	)
	username = this.GetString("user[username]")
	password = this.GetString("user[password]")
	fmt.Println(username, " : ", password)
	if ok := user.VerifyUser(username, password); ok {
		user.LoginUser(this.Ctx, username, true)
	} else {
		this.Abort("Error")
	}
}

func (this *UserController) logout() {
	user.LogoutUser(this.Ctx)
}

func (this *UserController) verifyUser(code string) {
	username, flag := user.VerifyUserActiveCode(code)
	if flag == false {
		this.Abort("worng code")
	}
	user.LoginUser(this.Ctx, username, true)
}

func (this *UserController) verifyResetPwd(code string) string {
	username, flag := user.VerifyUserActiveCode(code)
	if flag == false {
		this.Abort("worng code")
	}
	return username
}

func (this *UserController) verifyResetPwdPost(code string) {
	username, flag := user.VerifyUserActiveCode(code)
	if flag == false {
		this.Abort("worng code")
	}
	if username != this.GetString("user[username]") {
		this.Abort("user and code no match")
	}
	pwd := this.GetString("user[password]")
	user.SaveNewPassword(username, pwd)
}
