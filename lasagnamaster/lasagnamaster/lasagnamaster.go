package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePeraLayer int) int {
	if timePeraLayer == 0 {
		timePeraLayer = 2
	}

	return len(layers) * timePeraLayer
}

// TODO: define the 'Quantities()' function
func Quantities(sauceNoodlesLayers []string) (noodles int, sauce float64) {
	countNoodlesLayers := 0
	countSauceLayers := 0

	for i := 0; i < len(sauceNoodlesLayers); i++ {
		if sauceNoodlesLayers[i] == "noodles" {
			countNoodlesLayers++
			continue
		} else if sauceNoodlesLayers[i] == "sauce" {
			countSauceLayers++
		}
	}

	noodles = countNoodlesLayers * 50
	sauce = float64(countSauceLayers) * 0.2

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := append([]float64(nil), quantities...)

	for i := 0; i < len(result); i++ {
		result[i] *= 0.5 * float64(portions)
	}

	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
