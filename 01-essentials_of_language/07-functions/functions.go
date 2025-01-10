package main

import "fmt"

func main() {
	outputText1("test 1", "test 2", 123)
	outputText2("test 2", "test 4")
	fmt.Println(outputText3("Julian", "banana 🍌"))
	fmt.Println(outputText4("Julian", "apple 🍎"))
	fmt.Println(outputText5("Julian", "corn 🌽"))
}

// in go we have to provide the types for the arguments for the function
func outputText1(text1 string, text2 string, number int) {
	fmt.Println(text1, text2, number)
}

// when all the types for the function are the same we can shorten the type definition
// (you cannot when using the shorter way have different types of arguments in the
// function)
func outputText2(text1, text2 string) {
	fmt.Println(text1, text2)
}

// we can also return a value from the function and provide the information about
// what type of value our function returns
func outputText3(name, prise string) string {
	result := fmt.Sprintf("\nCongrats %v. You won: %v", name, prise)
	return result
}

// for functions that return more than one value we have to add them in the
// parentheses
func outputText4(name, prise string) (string, string) {
	result1 := fmt.Sprintf("\nCongrats %v.", name)
	result2 := fmt.Sprintf("You won: %v", prise)
	return result1, result2
}

// you can also provide the names for the variables you want to return, this way
// you don't have to declare them later on, you can just simply override them with
// new value
func outputText5(name, prise string) (result1 string, result2 string) {
	result1 = fmt.Sprintf("\nCongrats %v.", name)
	result2 = fmt.Sprintf("You won: %v", prise)
	return result1, result2
	// alternately you can just do the "return" (without the variables) and the
	// return variables that were defined with the function declaration will get
	// returned from this function
}
