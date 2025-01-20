package main

import "fmt"

func main() {
	// using variadic function that can take any number of numbers (type int)
	fmt.Println(sumUp1(1, 10, 14))
	fmt.Println(sumUp2(101, 1, 10, 14))

	// you can also use the slice as an parameter you just have to split the slice
	// into separate numbers
	numbers := []int{1, 4, 7, 99}
	fmt.Println(sumUp1(numbers...))
}

// variadic function that can take any amount of parameters
func sumUp1(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

// we can also have other arguments in a variadic function (here first argument
// when calling the function will be "startingValue" and rest will be treated as
// "numbers")
func sumUp2(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}
	return sum
}
