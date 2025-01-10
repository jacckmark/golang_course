package main

import (
	"fmt"
	"math"
)

func main() {
	// we can declare the const variables (like in JS this is a constant that cannot
	// be reassigned like the var variable)
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Investment amount: ")
	// this line will allow us to scan for the value entered into terminal and will
	// assign this value to the provided variable (notice the ampersand symbol, this is
	// a pointer that points to this variable)
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
