package main

import "fmt"

// here we are creating custom type that refers to the build in type in golang
type str string

// this can be useful for example to attach the methods to the custom type that
// can be used later on (we cannot attach the method to string, it will result in
// an error)
func (text str) log() {
	fmt.Println(text)
}

func main() {
	// using custom type (type alias) in code
	var name str = "Max"
	name.log()
}
