package main

import (
	"fmt"
	"strconv"
	"strings"
)

var userSelection string
var userSelectionValue int

func userIntInput() int{
	fmt.Scanln(&userSelection)
	userSelectionValue, _ = strconv.Atoi(userSelection)
	return userSelectionValue
}

func backToMain() {
	fmt.Printf("Back to Main Menu? (y/n): ")
	fmt.Scanln(&userSelection)

	if userSelection == "y" {
		main()
	}else if userSelection == "n" {
		backToMain()
	}else {
		fmt.Println("Not a valid selection.")
		backToMain()
	}
}

func userStringInput() string{
	fmt.Scanln(&userSelection)
	userSelectionCapital := strings.Title(strings.ToLower(userSelection))
	return userSelectionCapital
}
