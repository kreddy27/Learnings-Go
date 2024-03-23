package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, errRevenue := getUserInput("Enter revenue: ")
	expenses, errExpenses := getUserInput("Enter expenses: ")
	taxRate, errTaxRate := getUserInput("Enter taxrate: ")

	if errRevenue != nil || errExpenses != nil || errTaxRate != nil {
		fmt.Println("Invalid input")

	}
	fmt.Println("Revenue:", revenue)

	ebt, profit := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Ebt:", ebt)
	fmt.Println("Profit:", profit)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64) {
	ebt := revenue - expenses
	profit := revenue - expenses*taxRate/100
	result := fmt.Sprintf("Ebt :%1f\nProfit:%1f", ebt, profit)
	writeOutputToFile(result)
	return ebt, profit
}

func getUserInput(text string) (float64, error) {
	var number float64
	fmt.Print(text)
	fmt.Scan(&number)
	if number <= 0 {
		return number, errors.New("invalid input")
	}
	return number, nil
}

func writeOutputToFile(result string) {
	fileName := "result.txt"
	os.WriteFile(fileName, []byte(result), 0644)
}
