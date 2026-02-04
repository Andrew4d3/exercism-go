package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}

	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	var noodlesTotal int
	var sauceTotal float64

	for _, value := range layers {
		if value == "noodles" {
			noodlesTotal += 50
		} else if value == "sauce" {
			sauceTotal += 0.2
		}
	}

	return noodlesTotal, sauceTotal
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendList, myList []string) {
	missingIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = missingIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, targetPortions int) []float64 {
	scaledPortions := make([]float64, 0, len(quantities))
	for _, portionFor2 := range quantities {
		portionFor1 := portionFor2 / 2
		scaledPortions = append(scaledPortions, portionFor1*float64(targetPortions))
	}

	return scaledPortions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
