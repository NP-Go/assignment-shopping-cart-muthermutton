package main

func main() {

	showMainMenu()
	userMenuSel()

	switch userSelectionValue {
	case 1:
		printWholeList()
		backToMain()
	case 2:
		showGenReportMenu()
		userMenuSel()

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
		
	}

	// fmt.Printf("Category: %v - Item: %v, Quantity: %v, Cost: %v", category[shoppingList["Cake"].category], )
}
