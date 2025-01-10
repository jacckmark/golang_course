package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

// we return here not only the data but error if one occurs
func getBalanceFromFile() (float64, error) {
	// here we can get the error when the file does not exist for example
	data, err := os.ReadFile(accountBalanceFile)

	// to handle errors in the golang we check whether or not the error is not nil
	// we could also check if the data variable is nil but it is better to check the
	// error and then proceed with error handling
	if err != nil {
		return 1000, errors.New("failed to find the balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	// we need to add here nil because we might return an error object with the balance
	// (which probably in case of error will be nil). This is because of the function
	// definition which tells us we will return 2 variables
	return balance, nil
}

func main() {
	// we might get an error passed from the getBalanceFromFile method so we should
	// handle this case here
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// we can return from this function and it would stop the program execution,
		// but we can also use the panic method. This method will stop the execution,
		// output the error and print the stacktrace
		panic("Can't continue, sorry")
	}

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
