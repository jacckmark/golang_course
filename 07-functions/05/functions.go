package main

import "fmt"

func main() {
	fact1 := factorial1(5)
	fmt.Println(fact1)
	fact2 := factorial2(5)
	fmt.Println(fact2)
}

// factorial written with loop
func factorial1(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result = result * i
	}

	return result
}

// factorial written with recursion
func factorial2(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial2(num-1)
}
