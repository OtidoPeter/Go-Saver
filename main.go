package main

import (
	"errors"
	"fmt"
	"gosaver/example.gosaver/fileops"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"
const currentBudgetFile = "budget.txt"

func getBudgetFromFile() (float64, error) {
	data, err := os.ReadFile(currentBudgetFile)
	if err != nil {
		return 0, errors.New("Failed to find budget file")
	}
	budgetText := string(data)
	budget, err := strconv.ParseFloat(budgetText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse budget value to float")
	}

	return budget, nil
}

func writeBudgetToFile(budget float64) {
	budgetText := fmt.Sprint(budget)
	os.WriteFile(currentBudgetFile, []byte(budgetText), 0644)
}

func main() {
	const separator = "-----------------------"
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
	}
	var setBudget float64
	setBudget, err = getBudgetFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
	}
	fmt.Println("Welcome to GoSaver v1.0")
	fmt.Println(separator)

	for {

		operations()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your set budget: ")
			fmt.Scan(&setBudget)

			if setBudget <= 0 {
				fmt.Println("Invalid amount! Should be greater than 0")
				//return
				continue
			}
			if setBudget > accountBalance {
				fmt.Println("Invalid amount! Not enough money in your bank account")
				//return
				continue
			}

			fmt.Println("Your set budget:", setBudget)
			writeBudgetToFile(setBudget)
		case 4:
			fmt.Println("Your account balance:", accountBalance)
			fmt.Println("Your set budget:", setBudget)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}

}
