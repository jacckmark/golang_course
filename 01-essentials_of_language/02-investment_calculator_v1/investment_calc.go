package main

import (
	"fmt"
	"math"
)

func main() {
	// go is statically types language: int, float64, string, bool etc. We cannot mix and
	// match the types like in JS. We have to convert the values to use them together.
	// When we don't declare the type of value it is inferred by the compiler (we should
	// when possible declare the types because they can be inferred incorrectly)
	var investmentAmount float64
	// the investmentAmount variable has value of null if we didn't explicitly set the
	// value (like here). The null value depends on the type. For example:
	// int => 0
	// float64 => 0.0
	// string => ""
	// bool => false
	investmentAmount = 1000
	// if we don't declare the type (automatically inferred variables) for variable
	// we can use short assignment (this is recommended way for non-typed variables)
	// this sadly will not work if we declare the variables outside of the function
	// definition, then we have to use the longer notation
	expectedReturnRate := 5.5
	var years float64 = 10

	// we can declare variables in one line, but if we provide the type all of the
	// variables should have the same type
	var num1, num2 float64 = 100, 200

	// type conversion for "years" because it is inferred as int and we cannot do calculations
	// in go on two types simultaneously
	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println(futureValue)
	fmt.Println(num1 + num2)
}
