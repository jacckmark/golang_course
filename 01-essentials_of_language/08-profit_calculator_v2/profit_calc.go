package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserRevenue()
	expenses = getUserExpenses()
	taxRate = getUserTaxRate()

	earningsBeforeTax := calculateEarningsBeforeTax(revenue, expenses)
	earningsAfterTax := calculateEarningsAfterTax(earningsBeforeTax, taxRate)
	ratio := calculateRatio(earningsBeforeTax, earningsAfterTax)

	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Printf("%.1f\n", earningsAfterTax)
	fmt.Printf("%.3f\n", ratio)
}

func calculateEarningsBeforeTax(revenue, expenses float64) float64 {
	return revenue - expenses
}

func calculateEarningsAfterTax(earningsBeforeTax, taxRate float64) float64 {
	return earningsBeforeTax * (1 - taxRate/100)
}

func calculateRatio(earningsBeforeTax, earningsAfterTax float64) float64 {
	return earningsBeforeTax / earningsAfterTax
}

func getUserRevenue() float64 {
	var revenue float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	return revenue
}

func getUserExpenses() float64 {
	var expenses float64
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	return expenses
}

func getUserTaxRate() float64 {
	var taxRate float64
	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)
	return taxRate
}
