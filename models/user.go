package models

import "golang.org/x/crypto/bcrypt"

//created models here (how data will be stored)
type User struct {
	Id        uint   `json : "id"`
	FirstName string `json : first_name`
	LastName  string `json : last_name`
	EmailId   string `json : email`
	Password  []byte `json :"-"`
	Phone     string `json : phone`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14) //using _ for blank identifier as this function returns 2 values
	user.Password = hashedPassword

}
