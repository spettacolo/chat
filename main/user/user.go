package user

import "reflect"

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

func NewUser(id int, username, firstName, lastName, profilePicture, email, password string) *User {
	return &User{
		ID:             id,
		Username:       username,
		FirstName:      firstName,
		LastName:       lastName,
		ProfilePicture: profilePicture,
		Email:          email,
		Password:       password,
	}
}

func (u *User) GetUser() *User {
	return u
}

func (u *User) UpdateUser(fields map[string]interface{}) {
	val := reflect.ValueOf(u).Elem()
	for key, value := range fields {
		field := val.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}
}

func (u *User) DeleteUser() {
	u = nil
}
