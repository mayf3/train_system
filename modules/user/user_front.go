package user

import (
	"strings"

	"train_system/models"
	//"train_system/setting"
)

// CanRegistered checks if the username or e-mail is available.
func CanRegistered(username string, email string) (b1 bool, b2 bool) {
	b1 = models.HasUserByUsername(username)
	b2 = models.HasUserByEmail(email)
	return b1, b2
}

// check if exist user by username or email
func HasUser(username string) bool {
	if strings.IndexRune(username, '@') == -1 {
		return models.HasUserByUsername(username)
	}
	return models.HasUserByEmail(username)
}
