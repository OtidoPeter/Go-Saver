package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func operations() {
	const separator = "------------------------"
	fmt.Println("You can reach us 24/7 on", randomdata.PhoneNumber())
	fmt.Println(separator)

	fmt.Println("What do you want to do?")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Set Budget")
	fmt.Println("4. Check Balance and Budget Status")
	fmt.Println("5. Exit")
}
