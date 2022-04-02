package main

import (
	"fmt"
	"strconv"
	"strings"
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
		fmt.Println("Not a valid selection.")
		backToMain()
	}
}

func userStingValidation(){
	fmt.Scanln(&userSelection)
	userSelection = strings.Title(strings.ToLower(userSelection))
}
