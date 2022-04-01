package main

type itemInfo struct {
	category int
	quantity int
	unitCost float64
}

var category = []string{"Household", "Food", "Drinks"}
var shoppingList map[string]itemInfo

func init() {
	shoppingList = make(map[string]itemInfo)
	shoppingList["Fork"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	shoppingList["Plates"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	shoppingList["Cups"] = itemInfo{category: 0, quantity: 5, unitCost: 3}
	shoppingList["Bread"] = itemInfo{category: 1, quantity: 2, unitCost: 2}
	shoppingList["Cake"] = itemInfo{category: 1, quantity: 3, unitCost: 1}
	shoppingList["Coke"] = itemInfo{category: 2, quantity: 5, unitCost: 2}
	shoppingList["Sprite"] = itemInfo{category: 2, quantity: 5, unitCost: 2}
}
