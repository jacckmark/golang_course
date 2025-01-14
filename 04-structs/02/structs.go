package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// structs can also have the methods attached to them, so the method would be a
// function that gets appended into struct and is callable (notice the change in
// parenthesis placement)
func (u user) outputUserDetails() {
	// to turn this into an struct method we add the parenthesis before the function
	// name and place there the arguments (the whole parenthesis is called a receiver)
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// this method modifies the data in the struct (notice the pointer, if omitted
// we will not be able to edit the value in the struct because we are passing
// around an copy of the struct)
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser = user{firstName: userFirstName, lastName: userLastName, birthDate: userBirthDate, createdAt: time.Now()}
	// to call the attached method (attached to the struct user) we need to use dot
	// notation (notice that we don't need no arguments there because receiver gets
	// them from the struct)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
