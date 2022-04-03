package main

import (
	userInput "assignment-shopping-cart-muthermutton/userInput"
	"fmt"
)

func main() {

	showMainMenu()
	userSelectionValue := userInput.UserIntInput("\nselect your choice: ")
	

	switch userSelectionValue {
	case 1:
		printWholeList()
		backToMain()
	case 2:
		showGenReportMenu()
		userSelectionValue = userInput.UserIntInput("\nChoose your report or exit to Main Menu: ")
		switch userSelectionValue {
		case 1:
			totalCostCategory()
			backToMain()
		case 2:
			listByCategory()
			backToMain()
		case 3:
			main()
		}
	case 3:
		addItemFeature()
		main()
	case 4:
		modifyItem()
		backToMain()
	case 5:
		deleteItem()
		backToMain()
	case 6:
		printCurrentData()
		backToMain()
	case 7:
		addNewCategory()
		backToMain()
	case 8:
		break
	default:
		fmt.Println("Not a valid option!")
		main()
	}
}
