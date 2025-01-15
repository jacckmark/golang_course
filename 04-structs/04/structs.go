package main

import (
	"fmt"

	"example.com/test/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// using struct type to define the type we are expecting
	var appUser *user.User
	// creating new user using the constructor from the external package
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// here using the struct AdminUser constructor to create new struct
	admin := user.NewAdmin("test@gmail.com", "test123")
	// this method is available because it is a method that is a part of User struct
	// and we set the embedded User field as public one (we cannot however access the
	// fields like email and password because they are private (written with lowercase))
	admin.OutputUserDetails()
	// we can also access the methods that were embedded into AdminUser from User
	// using the notation and moving into embedded struct (but it is not necessary
	// here)
	// admin.User.OutputUserDetails()

	// here using user struct methods that were exposed to make them available in
	// other packages
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
