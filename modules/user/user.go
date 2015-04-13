package user

import (
	"fmt"
	"strings"

	"train_system/models"
	"train_system/modules/utils"
	"train_system/setting"
)

func GeneratePassword(password, salt string) string {
	password = utils.EncodePassword(password, salt, setting.PasswordLen)
	password = fmt.Sprintf("%s$%s", salt, password)
	return password
}

// register create user
func RegisterUser(username, email, password string) error {
	var (
		err  error
		salt string
		user models.User
	)
	user.Username = strings.ToLower(username)
	user.Email = strings.ToLower(email)
	user.Nickname = user.Username
	// use username as default nickname.
	// use random salt encode password
	salt = utils.GetRandomString(setting.SaltLen)
	user.Password = GeneratePassword(password, salt)
	user.Power = "user"
	err = user.Insert()
	return err
}

// set a new password to user
func SaveNewPassword(username, password string) error {
	var (
		salt string
	)
	salt = utils.GetRandomString(setting.SaltLen)
	password = GeneratePassword(password, salt)
	return models.UpdatePassword(username, password)
}
