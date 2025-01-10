package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// to concatenate the string you just need to provide additional arguments to
	// the println function
	fmt.Println("Future value: ", futureValue)
	// another way to concatenate string, you can use the printf and the %v will be
	// replaced by the second argument passed to the function
	fmt.Printf("Future value (adjusted for inflation): %v", futureRealValue)
	// for more than one string variable you just provide more arguments and more %v
	// in your initial string
	fmt.Printf("\nFuture value: %v\nFuture value (adjusted for inflation): %v", futureValue, futureRealValue)
	// to format provided values (here shortening the decimal places to one) we use
	// the %f with .2. The number controls how many of the decimal places to show
	// (there are more parameters that help you with string formatting in the official
	// docs)
	fmt.Printf("\nFuture value: %.2f\nFuture value (adjusted for inflation): %.2f", futureValue, futureRealValue)

	// you can also prepare the formatted string beforehand store it into variable
	// and use later
	formattedFV := fmt.Sprintf("\nCongrats \nFuture value: %.4f", futureValue)
	fmt.Println(formattedFV)

	// you can also use the backtick notation to break the line more naturally
	// (without use of \n)
	formattedFV2 := fmt.Sprintf(`
	Congrats 
	Future value: %.4f`, futureValue)
	fmt.Print(formattedFV2)
}
