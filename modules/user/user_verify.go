package user

import (
	"encoding/hex"

	"train_system/models"
	"train_system/modules/utils"
	"train_system/setting"
)

// verify username/email and password
func VerifyUser(username, password string) bool {
	right_pwd := models.GetPasswordByUsername(username)
	if VerifyPassword(password, right_pwd) {
		return true
	}
	return false
}

// compare raw password and encoded password
func VerifyPassword(password, right_pwd string) bool {
	salt := right_pwd[:setting.SaltLen]
	password = GeneratePassword(password, salt)
	return password == right_pwd
}

// get user by erify code
func getVerifyUser(user *models.User, code string) bool {
	if len(code) <= utils.TimeLimitCodeLength {
		return false
	}
	// use tail hex username query user
	hexStr := code[utils.TimeLimitCodeLength:]
	if b, err := hex.DecodeString(hexStr); err == nil {
		err := user.GetUserByUsername(string(b))
		if err == nil {
			return true
		}
	}
	return false
}

// verify active code when active account
func VerifyUserActiveCode(code string) (string, bool) {
	var user models.User
	minutes := setting.ActiveCodeLives
	if getVerifyUser(&user, code) {
		// time limit code
		prefix := code[:utils.TimeLimitCodeLength]
		data := utils.ToStr(user.Id) + user.Email + user.Username + user.Password + user.Rands
		if utils.VerifyTimeLimitCode(data, minutes, prefix) {
			user.SetVerify()
			return user.Username, true
		}
	}
	return "", false
}

// create a time limit code for user active
func CreateUserActiveCode(username string, startInf interface{}) string {
	var user models.User
	user.GetUserByUsername(username)
	minutes := setting.ActiveCodeLives
	data := utils.ToStr(user.Id) + user.Email + user.Username + user.Password + user.Rands
	code := utils.CreateTimeLimitCode(data, minutes, startInf)
	// add tail hex username
	code += hex.EncodeToString([]byte(user.Username))
	return code
}

// verify code when reset password
func VerifyUserResetPwdCode(code string) (string, bool) {
	var user models.User
	minutes := setting.ResetPwdCodeLives
	if getVerifyUser(&user, code) {
		// time limit code
		prefix := code[:utils.TimeLimitCodeLength]
		data := utils.ToStr(user.Id) + user.Email + user.Username + user.Password + user.Rands // + user.Updated.String()
		if utils.VerifyTimeLimitCode(data, minutes, prefix) {
			return user.Username, true
		}
	}
	return "", false
}

// create a time limit code for user reset password
func CreateUserResetPwdCode(username string, startInf interface{}) string {
	var user models.User
	user.GetUserByUsername(username)
	minutes := setting.ResetPwdCodeLives
	data := utils.ToStr(user.Id) + user.Email + user.Username + user.Password + user.Rands // + user.Updated.String()
	code := utils.CreateTimeLimitCode(data, minutes, startInf)
	// add tail hex username
	code += hex.EncodeToString([]byte(user.Username))
	return code
}
