package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	// for loop is the only loop in golang, because of it it can run fixed amount of
	// times like for in JS or work like while loop in JS and end when we use "break"
	// keyword
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

		if choice == 1 {
			fmt.Println("Your account balance is ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Printf("Balance updated! You deposited %v$. Your new balance is: %v$ \n", depositAmount, accountBalance)
		} else if choice == 3 {
			if accountBalance == 0 {
				fmt.Print("No money to withdraw. Good bye!")
				continue
			}
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			accountBalance -= withdrawalAmount
			fmt.Printf("Balance updated! You withdrew %v$. Your new balance is: %v$ \n", withdrawalAmount, accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")
}
