package models

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
	CreateAt int
	IsActive bool
}

//get a user object (not insert into db in this function)
func NewUser(info map[string]interface{}) *User {
	return &User{
		Username: info["username"].(string),
		Nickname: info["nickname"].(string),
		Email:    info["email"].(string),
	}
}

// default add a user and his related profile into database
func InsertUser(
	username string,
	password string,
	nickname string,
	email string) (*User, error) {

	da := DbAgent{}
	info := map[string]interface{}{
		"username":  username,
		"password":  password,
		"email":     email,
		"create_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	result, err := da.From("user").Insert(info)

	if err != nil {
		return nil, err
	}

	user := NewUser(info)
	user.Id, _ = result.LastInsertId()
	return user, nil
}

// this will delete profile cascade
func DeleteUser(username string) error {
	da := DbAgent{}
	_, err := da.From("user").Where("username = ?", username).Delete()
	return err
}
