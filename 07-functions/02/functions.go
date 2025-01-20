package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(doubleNumbers(&numbers))
	fmt.Println(universalNumbersMultiplier(&numbers, double))
	fmt.Println(universalNumbersMultiplier(&numbers, triple))
	// using function that returns function value to us ("getTransformerFunc")
	fmt.Println(universalNumbersMultiplier(&numbers, getTransformerFunc(&numbers)))
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

// we not only pass the function as arguments, we can also return them as values
// from the functions
func getTransformerFunc(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func universalNumbersMultiplier(numbers *[]int, transform func(int) int) []int {
	resNumbers := []int{}
	for _, val := range *numbers {
		resNumbers = append(resNumbers, transform(val))
	}
	return resNumbers
}
