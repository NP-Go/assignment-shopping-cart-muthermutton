package main

import (
	userInput "assignment-shopping-cart-muthermutton/userInput"
	"fmt"
)

func backToMain() {
	input := userInput.UserInputYN("Back to main menu?")

	if input {
		main()
	} else {
		backToMain()
	}
}

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
	fmt.Println("8. Modify Category Name")
	fmt.Println("9. Exit Program")
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

	newItemName = userInput.UserStringInput("What is the name of your Item?")

	//user input validation check for valid category
	for ok := true; ok; ok = invalidUserInput {

		newCategory = userInput.UserStringInput("What category does it belong to?")
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

	newQuantity = userInput.UserIntInput("How many units are there?")

	newUnitCost = float64(userInput.UserIntInput("How much does it cost per unit?"))

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

	for {
		// fmt.Println("\nWhich item would you wish to modify?")
		// userInput := userStringInput()
		userInput := userInput.UserStringInput("\nWhich item would you wish to modify?")

		mapValue, ok := shoppingList[userInput]

		if ok {
			itemToChange = mapValue
			originalItemName = userInput
			fmt.Printf("\nCurrent Item Name: %v - Category: %v - Quantity: %v - Unit Cost %v\n", userInput, categories[mapValue.category], mapValue.quantity, mapValue.unitCost)
			break
		} else {
			println("No such Item in shopping list")
		}
	}

	//collection of data to change
	newItemName = userInput.UserStringInput("Enter new Item Name. Enter for no change.")

	//Used different type of Do while loop
	invalidUserInput = true
	for {
		newCategory = userInput.UserStringInput("Enter new Category. Enter for no change.")
		if newCategory == "" {
			break
		}
		for catId, category := range categories {
			if newCategory == category {
				newCategoryIndex = catId
				invalidUserInput = false
			}
		}
		if !invalidUserInput {
			break
		}
		userReply := userInput.UserInputYN("No such category! Add new category?")
		if userReply {
			categories = append(categories, newCategory)
			newCategoryIndex = len(categories) - 1
			fmt.Println("New Category:", newCategory, "added at index", len(categories)-1)
			break
		}
	}

	newQuantity = userInput.UserIntInput("Enter new Quantity. Enter for no change.")

	newUnitCost = float64(userInput.UserIntInput("Enter new Unit Cost. Enter for no change."))

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
		fmt.Println("\nDelete Item.")
		itemNameDelete := userInput.UserStringInput("Enter item name to delete:")

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
	for {
		var categoryExists = false
		newCategory := userInput.UserStringInput("What is the New Category Name to add?")
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

func modifyCategory() {
	var changeAtIndex int
	var invalidUserInput = true
	var originalCategory string

	for ok := true; ok; ok = invalidUserInput {
		newCategory := userInput.UserStringInput("What category does it belong to?")
		for catId, category := range categories {
			if newCategory == category {
				changeAtIndex = catId
				originalCategory = category
				invalidUserInput = false
			}
		}
		if invalidUserInput {
			userReply := userInput.UserInputYN("No valid category. Exit to Main Menu?")
			if userReply {
				main()
			}
		}
	}
	modifyTo := userInput.UserStringInput("New Category Name: ")
	categories[changeAtIndex] = modifyTo
	fmt.Println("\nYou changed Category Name", originalCategory, "to", modifyTo)
}
