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

func genReportMenu() {
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category")
	fmt.Println("2. List of items by category")
	fmt.Println("3. Main Menu")
	fmt.Println("Choose your report or exit to Main Menu:")
}
