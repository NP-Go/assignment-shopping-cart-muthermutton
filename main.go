package main

func main() {

	showMainMenu()
	userMenuSel()

	switch userSelectionValue {
	case 1:
		printWholeList()
		backToMain()
	}

	// fmt.Printf("Category: %v - Item: %v, Quantity: %v, Cost: %v", category[shoppingList["Cake"].category], )
}
