package main

import (
	"fmt"
	"time"
)

// this is a struct that is working like an type definition that you reuse
type user struct {
	// here we are defining the type fields with their type
	firstName string
	lastName  string
	birthDate string
	// like regular types in JS you can nest the types in golang or use other types
	// (here using the type from the	time library)
	createdAt time.Time
}

// fields in struct are not required so you can pass them when creating an variable
// based on struct but you can omit them (all or just one for example)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// defining the variable based on user struct
	var appUser user
	// to assign the values to the variable when using the struct we use the struct
	// name followed by curly brackets where we pass the fields values
	appUser = user{firstName: userFirstName, lastName: userLastName, birthDate: userBirthDate, createdAt: time.Now()}
	// you can also pass the values without specifying which value is which but you
	// have to keep in mind that the order of variables in such variable needs to
	// match the struct definition (especially here where you have 3 strings in the
	// row and can accidentally pass the lastName as birthDate)
	// appUser = user{userFirstName, userLastName,userBirthDate, time.Now()}

	// you can also create an empty struct (when you don't pass the data right away)
	// the values will be null (so for float fields it will be 0.0 and for string
	// fields "")
	// appUser = user{}

	outputUserDetails(appUser)
}

// here defining that we expect the struct of type user
func outputUserDetails(u user) {
	// to access the struct properties we just access them using the dot notation
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
