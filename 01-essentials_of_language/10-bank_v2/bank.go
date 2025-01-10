package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// this function will take the current balance and write it into the 'balance.txt'
// file located in the current directory
func writeBalanceToFile(balance float64) {
	// we have to convert the float into string
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	// this function returns two variables data and the error, because we don't need
	// the error we use the underscore to tell go that we are aware it contains more
	// than one results but we don't need it
	data, _ := os.ReadFile(accountBalanceFile)
	// to convert byte data to string we can use build in 'string' method
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is ", accountBalance, "$")
		case 2:
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Printf("Balance updated! You deposited %v$. Your new balance is: %v$ \n", depositAmount, accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			if accountBalance == 0 {
				fmt.Print("No money to withdraw. Good bye!")
				continue
			}

			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Print("Invalid amount. You cannot withdraw more than you have!")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Printf("Balance updated! You withdrew %v$. Your new balance is: %v$ \n", withdrawalAmount, accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
