package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

// this is a generic type, here we say that T can be of any of given types (string,
// float64, or int). This also influences the return type because there we also use
// T as a type
func add[T int | float64 | string](a, b T) T {
	return a + b
}
