package main

func main() {


	showMainMenu()
	userIntInput()

	switch userSelectionValue {
	case 1:
		printWholeList()
		backToMain()
	case 2:
		showGenReportMenu()
		userIntInput()

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
	}
	

}
