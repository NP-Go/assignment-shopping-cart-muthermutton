package main

type itemInfo struct {
	category int
	quantity int
	unitCost float64
}

var categories = []string{"Household", "Food", "Drinks"}
var shoppingList = make(map[string]itemInfo)

func addNewItem(key string, newItemInfo itemInfo) {
	shoppingList[key] = newItemInfo
}

func init() {
	shoppingList["Fork"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	shoppingList["Plates"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	shoppingList["Cups"] = itemInfo{category: 0, quantity: 5, unitCost: 3}
	shoppingList["Bread"] = itemInfo{category: 1, quantity: 2, unitCost: 2}
	shoppingList["Cake"] = itemInfo{category: 1, quantity: 3, unitCost: 1}
	shoppingList["Coke"] = itemInfo{category: 2, quantity: 5, unitCost: 2}
	shoppingList["Sprite"] = itemInfo{category: 2, quantity: 5, unitCost: 2}
}
