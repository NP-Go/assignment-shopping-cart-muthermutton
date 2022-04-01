package main

import (
	"fmt"
	"strconv"
)

var userSelection string
var userSelectionValue int

func userMenuSel() {
	fmt.Scanln(&userSelection)
	userSelectionValue, _ = strconv.Atoi(userSelection)
}

func backToMain() {
	fmt.Printf("Back to Main Menu? (y/n): ")
	fmt.Scanln(&userSelection)

	if userSelection == "y" {
		main()
	}else if userSelection == "n" {
		backToMain()
	}else {
		println("Not a valid selection.")
		backToMain()
	}
}
