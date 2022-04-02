package main

import "fmt"

//printing and storing of different menus
func showMainMenu() {
	fmt.Println("\nShopping List Application")
	fmt.Println("=========================")
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate shopping List Report")
	fmt.Println("3. Add Items")
	fmt.Println("4. Modify Items")
	fmt.Println("5. Delete Item.")
	fmt.Println("6. Print Current Data")
	fmt.Println("7. Add New Category Name")
	fmt.Printf("\nSelect your choice: ")
}

// iterates over the map and prints list in unordered
func printWholeList() {
	fmt.Println("\nShopping List Contents: ")
	for itemName, item := range shoppingList {
		fmt.Printf("Category: %v - Item: %v, Quantity: %v, Cost: %v\n", categories[item.category], itemName, item.quantity, item.unitCost)
	}
	fmt.Println(" ")
}

func showGenReportMenu() {
	fmt.Println("\nGenerate Report")
	fmt.Println("1. Total Cost of each category")
	fmt.Println("2. List of items by category")
	fmt.Println("3. Main Menu")
	fmt.Printf("\nChoose your report or exit to Main Menu: ")
}

func totalCostCategory() {
	fmt.Println("\nTotal cost by Category.")
	for catId, category := range categories {
		var total float64
		for _, item := range shoppingList {
			if item.category == catId {
				total += (float64(item.quantity) * item.unitCost)
			}
		}
		fmt.Printf("%v cost : %v\n", category, total)
	}
	fmt.Println(" ")
}

func listByCategory() {
	fmt.Println(" ")
	for catId, category := range categories {
		for itemName, item := range shoppingList {
			if item.category == catId {
				fmt.Printf("Category: %v - Item: %v, Quantity: %v, Cost: %v\n", category, itemName, item.quantity, item.unitCost)
			}
		}
	}
	fmt.Println(" ")
}

func addItem() {
	var newItemName string
	var newCategory string
	var newCategoryIndex int
	var newQuantity int
	var newUnitCost float64
	var invalidUserInput bool = true

	fmt.Println("What is the name of your Item?")
	newItemName = userStringInput()

	//user input validation check for valid category
	for ok := true; ok; ok = invalidUserInput {
		fmt.Println("What category does it belong to?")
		newCategory = userStringInput()
		for catId, category := range categories {
			if newCategory == category {
				invalidUserInput = false
				newCategoryIndex = catId
			}
		}
		if invalidUserInput {
			println("No such category")
		}
	}

	fmt.Println("How many units are there?")
	newQuantity = userIntInput()

	fmt.Println("How much does it cost per unit?")
	newUnitCost = float64(userIntInput())

	shoppingList[newItemName] = itemInfo{category: newCategoryIndex, quantity: newQuantity, unitCost: newUnitCost}

}

func modifyItem() {
	var newItemName string
	// var newCategory string
	// var newQuantity int
	// var newUnitCost float64
	var invalidUserInput bool = true
	var stringInputValue string

	//user input validation check for valid category
	for ok := true; ok; ok = invalidUserInput {
		fmt.Println("\nWhich item would you wish to modify?")
		x := userStringInput()
		for itemName, itemInfo := range shoppingList {
			if x == itemName {
				fmt.Printf("\nCurrent Item Name: %v - Category: %v - Quantity: %v - Unit Cost %v\n", itemName, itemInfo.category, itemInfo.quantity, itemInfo.unitCost)
				invalidUserInput = false
			}
		}
		if invalidUserInput {
			println("No such Item in shopping list")
		}
	}

	fmt.Println("Enter new Item Name. Enter for no change.")
	stringInputValue = userStringInput()
	if stringinputvalue == " " {

	} else {

	}
}
