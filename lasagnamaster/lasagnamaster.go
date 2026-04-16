package main

import (
	l "exercism/lasagnamaster/lasagnamaster"
	"fmt"
)

var p = fmt.Println

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	p(l.PreparationTime(layers, 3))

	p(l.Quantities(layers))

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	l.AddSecretIngredient(friendsList, myList)
	p(myList)

	quantities := []float64{1.2, 3.6, 10.5}
	p(l.ScaleRecipe(quantities, 4))

}
