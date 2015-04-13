package base

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego"

	"train_system/models"
	"train_system/modules/user"
)

type BaseController struct {
	beego.Controller
	Username string
	IsLogin  bool
}

// Prepare : do something before
func (this *BaseController) Prepare() {
	var (
		flag bool
	)
	this.StartSession()

	this.IsLogin = false

	this.Username, flag = user.GetUserFromSession(this.CruSession)
	if flag == true {
		this.IsLogin = true
	} else {
		this.Username, flag = user.LoginUserFromRememberCookie(this.Ctx)
		if flag == true {
			this.IsLogin = true
		}
	}

	//	this.Username, flag = user.LoginUserFromRememberCookie(this.Ctx)
	//	if flag {
	//		this.IsLogin = true
	//	}

	this.Data["Permission"] = "anonymous"

	if this.IsLogin {
		this.Data["username"] = this.Username
		this.Data["permission"] = models.GetPowerByUsername(this.Username)
	}
	fmt.Println(this.CruSession)
	fmt.Println(this.Ctx.Input.Request.Cookies())
}

// Finish : do something after
func (this *BaseController) Finish() {
}

// check if not login then redirect
func (this *BaseController) CheckLoginRedirect(args ...interface{}) bool {
	var (
		code        int
		need_login  bool
		redirect_to string
	)
	code = 302
	need_login = true
	for _, arg := range args {
		switch v := arg.(type) {
		case bool:
			need_login = v
		case string:
			// custom redirect url
			redirect_to = v
		case int:
			// custom redirect url
			code = v
		}
	}
	// if need login then redirect
	if need_login && !this.IsLogin {
		if len(redirect_to) == 0 {
			req := this.Ctx.Request
			scheme := "http"
			if req.TLS != nil {
				scheme += "s"
			}
			redirect_to = fmt.Sprintf("%s://%s%s", scheme, req.Host, req.RequestURI)
		}
		redirect_to = "/login?to=" + url.QueryEscape(redirect_to)
		this.Redirect(redirect_to, code)
		return true
	}
	// if not need login then redirect
	if !need_login && this.IsLogin {
		if len(redirect_to) == 0 {
			redirect_to = "/"
		}
		this.Redirect(redirect_to, code)
		return true
	}
	return false
}
