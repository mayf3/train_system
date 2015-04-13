package user

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"

	"train_system/models"
	"train_system/modules/utils"
	"train_system/setting"
)

// login user
func LoginUser(ctx *context.Context, username string, remember bool) {
	// werid way of beego session regenerate id...
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateId(ctx.ResponseWriter, ctx.Request)
	ctx.Input.CruSession.Set("username", username)
	if remember {
		WriteRememberCookie(username, ctx)
	}
}

func WriteRememberCookie(username string, ctx *context.Context) {
	rand := models.GetRandsByUsername(username)
	pwd := models.GetPasswordByUsername(username)
	secret := utils.EncodeMd5(rand + pwd)
	days := 86400 * setting.LoginRememberDays
	ctx.SetCookie(setting.CookieUsername, username, days)
	ctx.SetSecureCookie(secret, setting.CookieRememberName, username, days)
}

func DeleteRememberCookie(ctx *context.Context) {
	ctx.SetCookie(setting.CookieUsername, "", -1)
	ctx.SetCookie(setting.CookieRememberName, "", -1)
}

func LoginUserFromRememberCookie(ctx *context.Context) (username string, success bool) {
	tmp_username := ctx.GetCookie(setting.CookieUsername)
	fmt.Println(tmp_username)
	if len(tmp_username) == 0 {
		return "", false
	}
	defer func() {
		if !success {
			DeleteRememberCookie(ctx)
		}
	}()
	if err := models.HasUserByUsername(tmp_username); err == false {
		fmt.Println("No username")
		return "", false
	}
	rand := models.GetRandsByUsername(tmp_username)
	pwd := models.GetPasswordByUsername(tmp_username)
	secret := utils.EncodeMd5(rand + pwd)
	value, _ := ctx.GetSecureCookie(secret, setting.CookieRememberName)
	if value != tmp_username {
		fmt.Println("Don't equal")
		return "", false
	}
	LoginUser(ctx, tmp_username, true)
	return tmp_username, true
}

// logout user
func LogoutUser(ctx *context.Context) {
	DeleteRememberCookie(ctx)
	ctx.Input.CruSession.Delete("username")
	ctx.Input.CruSession.Flush()
	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

// get user if key exist in session
func GetUserFromSession(sess session.SessionStore) (string, bool) {
	username := sess.Get("username")
	if username == nil {
		return "", false
	}
	return username.(string), true
}
