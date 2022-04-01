package main

type itemInfo struct {
	category int
	quantity int
	unitCost float64
}

var category = []string{"Household", "Food", "Drinks"}

var shoppingList map[string]itemInfo
