package main

import "fmt"

// functions is go are first class values (function can take variables, values,
// but they can also expect a function like in JS)
func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// notice that we expect an pointer in the function so we have to get the address
	// of this array
	fmt.Println(doubleNumbers(&numbers))

	// here we are using the function that is capable of taking a function as an
	// argument (notice a lack of parenthesis next to the function name that is an
	// argument, this is because we try to pass here function as a value (like variable)
	// and not try to run it)
	fmt.Println(universalNumbersMultiplier(&numbers, double))
	fmt.Println(universalNumbersMultiplier(&numbers, triple))
}

// this function takes numbers array and then doubles all the items in this array
// (1 becomes 2, 2 becomes 4 and so on)
func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// this function takes data and a function as the arguments and uses passed function
// to modify the result of running it (notice the type for our function argument
// we set it to the function that takes an int and returns an int). For more
// complex function types it is good to use custom types (like below):
// type transformFn func(int) int
func universalNumbersMultiplier(numbers *[]int, transform func(int) int) []int {
	resNumbers := []int{}
	for _, val := range *numbers {
		resNumbers = append(resNumbers, transform(val))
	}
	return resNumbers
}
