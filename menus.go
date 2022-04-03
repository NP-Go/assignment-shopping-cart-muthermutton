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

func addItemFeature() {
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

	addNewItem(newItemName, itemInfo{category: newCategoryIndex, quantity: newQuantity, unitCost: newUnitCost})

}

func modifyItem() {
	var newCategory string
	var newItemName string
	var originalItemName string
	var newCategoryIndex int
	var newQuantity int
	var newUnitCost float64
	var invalidUserInput bool = true
	var itemToChange itemInfo

	//user input validation check for valid category
	for ok := true; ok; ok = invalidUserInput {
		fmt.Println("\nWhich item would you wish to modify?")
		x := userStringInput()
		for itemName, itemInfo := range shoppingList {
			if x == itemName {
				itemToChange = itemInfo
				originalItemName = itemName
				fmt.Printf("\nCurrent Item Name: %v - Category: %v - Quantity: %v - Unit Cost %v\n", itemName, categories[itemInfo.category], itemInfo.quantity, itemInfo.unitCost)
				invalidUserInput = false
			}
		}
		if invalidUserInput {
			println("No such Item in shopping list")
		}
	}
	//collection of data to change
	fmt.Println("Enter new Item Name. Enter for no change.")
	newItemName = userStringInput()

	//Used different type of Do while loop
	invalidUserInput = true
	for {
		fmt.Println("Enter new Category. Enter for no change.")
		newCategory = userStringInput()
		for catId, category := range categories {
			if newCategory == category {
				newCategoryIndex = catId
				invalidUserInput = false
			}
		}
		if !invalidUserInput {
			break
		}
		fmt.Println("No such category!")
	}

	fmt.Println("Enter new Quantiy. Enter for no change.")
	newQuantity = userIntInput()

	fmt.Println("Enter new Unit Cost. Enter for no change.")
	newUnitCost = float64(userIntInput())

	if newCategory != "" {
		itemToChange.category = newCategoryIndex
	} else {
		fmt.Println("No changes made to Category.")
	}

	if newQuantity != 0 {
		itemToChange.quantity = newQuantity
	} else {
		fmt.Println("No changes made to quantity.")
	}

	if newUnitCost != 0 {
		itemToChange.unitCost = newUnitCost
	} else {
		fmt.Println("No changes made to cost.")
	}

	if newItemName != originalItemName {
		addNewItem(newItemName, itemInfo{category: newCategoryIndex, quantity: newQuantity, unitCost: newUnitCost})
		delete(shoppingList, originalItemName)
	} else {
		fmt.Println("No changes made to item name.")
	}

	// fmt.Printf(shoppingList)
}

func deleteItem() {
	for {
		fmt.Println("Delete Item.\n")
		itemNameDelete := userStringInput()

		_, exist := shoppingList[itemNameDelete]

		if exist {
			delete(shoppingList, itemNameDelete)
			fmt.Println("\nItem ", itemNameDelete, "has been deleted.")
			break
		} else {
			fmt.Println("\nItem does not exist. Please enter valid item.")
		}
	}
}

func printCurrentData() {
	fmt.Println("\nPrint Current Data.")
	for itemName, itemInfo := range shoppingList {
		fmt.Printf("%v - %v\n", itemName, itemInfo)
	}
	fmt.Println("")
}

func addNewCategory() {

	fmt.Println("Add New Category Name.")
	var categoryExists = false
	for {
		fmt.Println("What is the New Category Name to add?")
		newCategory := userStringInput()
		for catId, category := range categories {
			if newCategory == category {
				fmt.Println("Category", category, "already exist at index", catId)
				categoryExists = true
			}
		}
		if !categoryExists {
			categories = append(categories, newCategory)
			fmt.Println("New Category:", newCategory, "added at index", len(categories)-1)
			break
		}
	}
}
