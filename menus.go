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
	for k, item := range shoppingList {
		fmt.Printf("Category: %v - Item: %v, Quantity: %v, Cost: %v\n", category[shoppingList[k].category], k, item.quantity, item.unitCost)
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
	var total0 float64
	var total1 float64
	var total2 float64

	for _, item := range shoppingList {
		if item.category == 0 {
			total0 += (float64(item.quantity) * item.unitCost)
		} else if item.category == 1 {
			total1 += (float64(item.quantity) * item.unitCost)
		} else {
			total2 += (float64(item.quantity) * item.unitCost)
		}
	}

	fmt.Println("\nTotal cost by Category.")
	fmt.Println("Household cost :  ", total0)
	fmt.Println("Food cost :  ", total1)
	fmt.Println("Drinks cost :  ", total2,)
	fmt.Println("")
}
