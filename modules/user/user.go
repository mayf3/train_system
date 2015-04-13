// https://github.com/beego/wetalk
package user

import (
	"fmt"
	"strings"

	"train_system/models"
	"train_system/modules/utils"
	"train_system/setting"
)

// generate password by password && salt
func GeneratePassword(password, salt string) string {
	password = utils.EncodePassword(password, salt, setting.PasswordLen)
	password = fmt.Sprintf("%s$%s", salt, password)
	return password
}

// register create user
func RegisterUser(username, email, password string) error {
	var user models.User
	// use username as default nickname.
	user.Username = strings.ToLower(username)
	user.Email = strings.ToLower(email)
	user.Nickname = user.Username
	// use random salt encode password
	salt := models.GenerateSalt()
	user.Password = GeneratePassword(password, salt)
	user.Power = "user"
	err := user.Insert()
	return err
}

// set a new password to user
func SaveNewPassword(username, password string) error {
	salt := models.GenerateSalt()
	password = GeneratePassword(password, salt)
	return models.UpdatePassword(username, password)
}
