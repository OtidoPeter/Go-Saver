package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"
const currentBudgetFile = "budget.txt"

func writeBudgetToFile(budget float64) {
	budgetText := fmt.Sprint(budget)
	os.WriteFile(currentBudgetFile, []byte(budgetText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find account balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse account balance value to float")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	const separator = "-----------------------"
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
	}
	var setBudget float64
	fmt.Println("Welcome to GoSaver v1.0")
	fmt.Println(separator)

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Set Budget")
		fmt.Println("4. Check Balance and Budget Status")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var depositAmount float64
			fmt.Print("Your deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount! Should be greater than 0")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your updated account balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 2:
			var withdrawalAmount float64
			fmt.Print("Your withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount! Not enough money in your bank account")
				//return
				continue
			}
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount! Should be greater than 0")
				//return
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your updated account balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your set budget: ")
			fmt.Scan(&setBudget)

			if setBudget <= 0 {
				fmt.Println("Invalid amount! Should be greater than 0")
				//return
				continue
			}

			fmt.Println("Your set budget:", setBudget)
			writeBudgetToFile(setBudget)
		case 4:
			fmt.Println(accountBalance)
			fmt.Println(setBudget)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}
