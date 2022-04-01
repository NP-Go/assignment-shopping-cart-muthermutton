package main

import (
	"fmt"
	"strconv"
)

var userSelection string
var userSelectionValue int

func userMenuSel() {
	fmt.Scanln(&userSelection)
	userSelectionValue, _ = strconv.Atoi(userSelection)
}
