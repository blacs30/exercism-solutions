package lasagna

import (
	"fmt"
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	fmt.Println(layers, prepTime)
	if prepTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * prepTime

}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		} else {
			fmt.Println("Unkown layer")
		}
	}
	return
}

// AddSecretIngredient takes to []string with friends list of ingrediends and mine
// it compares the list and finds the secret ingredient of my friend and replaces
// the last element of my list (which is a ?) with the secret ingredeint
func AddSecretIngredient(friendsList, myList []string) {
	fmt.Println(friendsList, myList)
	var secretIngredient = ""
	for _, hisItem := range friendsList {
		found := false
		for _, myItem := range myList {
			if hisItem == myItem {
				found = true
				continue
			}
		}
		if !found {
			secretIngredient = hisItem
			fmt.Println(secretIngredient)
			myList[len(myList)-1] = secretIngredient
		}
	}
	return
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portion int) []float64 {
	fmt.Println(amount, portion)
	var new_values = []float64{}
	for _, item := range amount {
		new_amount := item / float64(2) * float64(portion)
		new_values = append(new_values, new_amount)
	}
	return new_values
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
