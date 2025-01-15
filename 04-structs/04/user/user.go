package user

import (
	"errors"
	"fmt"
	"time"
)

// to make this package accessible in other packages you need to make the names of
// functions uppercase, but for structs you need to also change the struct name
// ane struct field names to uppercase to make them available (in this case we
// only need the struct name available and not the fields)
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// here we are creating an struct that uses User struct (we embed it into AdminUser)
type AdminUser struct {
	email    string
	password string
	// you can give it a name other than "User" or just pass it as is and then this
	// field will be accessible under "User"
	User
}

// had to change the user struct to uppercase
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// for the constructors we often can observe the pattern that the name is just "New"
// this is the pattern that can be seen in external packages for example (look line
// 36 where we create the error using error package constructor)
func New(firstName, lastName, birthDate string) (*User, error) {
	fmt.Println("text")
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth data are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// a constructor method for AdminUser struct would have to have another name than
// "New", because this name is already taken
func NewAdmin(email, password string) AdminUser {
	return AdminUser{
		email:    email,
		password: password,
		// here providing embedded User struct
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}
