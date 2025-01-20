package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// here we use the anonymous function because maybe we don't use the multiplying
	// function more than once. In such a case we can define the function in place
	// without name and pass it as is
	fmt.Println(universalNumbersMultiplier(&numbers, func(num int) int { return num * 5 }))
}

func universalNumbersMultiplier(numbers *[]int, transform func(int) int) []int {
	resNumbers := []int{}
	for _, val := range *numbers {
		resNumbers = append(resNumbers, transform(val))
	}
	return resNumbers
}
