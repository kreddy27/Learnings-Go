/*
1. Check balance
2. Deposit money
3. Withdraw mone
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 100, errors.New("failed to read file")
	}
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}

func writeToFiles(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println(err)
		//panic("Sorry")
	}
	fmt.Println("Welcome to bank")
	for {
		//Prints welcome  message
		welcomeMessage()
		var choice int
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1

		if choice == 1 {
			fmt.Println("Your account balance is: ", accountBalance)
			writeToFiles(accountBalance)
		} else if choice == 2 {
			fmt.Println("Enter amount to deposit")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				return
			}
			accountBalance = accountBalance + depositAmount
			writeToFiles(accountBalance)

			fmt.Println("Your total balance is: ", accountBalance)
		} else if choice == 3 {
			fmt.Println("Enter amount to withdraw")
			var amountToWithdraw float64
			fmt.Scan(&amountToWithdraw)
			if amountToWithdraw > accountBalance {
				fmt.Println("You don't have enough balance to withdraw")
				//return
				continue
			}
			accountBalance = accountBalance - amountToWithdraw
			writeToFiles(accountBalance)

			fmt.Println("Your withdraw has done")
			fmt.Println("Your total balance is: ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
}

func welcomeMessage() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("Please enter your choice")
}
