package main

import "fmt"

func main() {
	age := 32

	// here we create a pointer to age variable (it points to the localization where
	// age is stored, this way we can for example mutate the age, by changing the
	// variable address). Type of a pointer is always written with asterix (here
	// *int). Uninitialized pointer has value of nil
	var agePointer *int
	agePointer = &age

	fmt.Println("Age variable memory address: ", agePointer)
	// to access the value stored where the pointer points you have to dereference
	// the pointer (you add the asterix before it)
	fmt.Println("Age variable value: ", *agePointer)

	adultYears := getAdultAge(agePointer)
	fmt.Println(adultYears)

	setAdultAge(agePointer, 13)
	fmt.Println("New age after pointer mutation: ", age)
}

// we can also pass the pointer into functions (we only have to specify the correct
// type by adding asterix)
func getAdultAge(agePointer *int) int {
	// here we are looking up the value stored under the passed pointer and then
	// subtracting the value (arithmetic operations done directly on pointers are
	// not possible in go)
	return *agePointer - 18
	// using pointers is here not necessary. We use them generally to avoid copying
	// variables over and over when we pass them to functions for example (or directly
	// mutating the variables). Here the variable is small and resigning from copying
	// it will not result in a big memory change (we would not even notice it). So
	// this is not a good example of using pointers to avoid copying variables in golang
}

func setAdultAge(agePointer *int, newAgeValue int) {
	// here we are using pointers to directly edit the variable value. We dereferencing
	// it and assigning the new value to it (we take the place in memory from the pointer
	// and put there in different data)
	*agePointer = newAgeValue
}
