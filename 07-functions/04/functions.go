package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubleFunction := createTransformer(2)
	quadrupleFunction := createTransformer(4)
	fmt.Println(universalNumbersMultiplier(&numbers, doubleFunction))
	fmt.Println(universalNumbersMultiplier(&numbers, quadrupleFunction))
}

func universalNumbersMultiplier(numbers *[]int, transform func(int) int) []int {
	resNumbers := []int{}
	for _, val := range *numbers {
		resNumbers = append(resNumbers, transform(val))
	}
	return resNumbers
}

// this is an example of closure. We have the access to the "factor" which is an
// argument of function "createTransformer" inside nested anonymous function and
// we can use it (because of closures)
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
