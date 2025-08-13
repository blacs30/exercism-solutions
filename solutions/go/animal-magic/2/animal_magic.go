package chance

import (
	"fmt"
	"math/rand/v2"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return 1 + rand.IntN(19)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	randInt := rand.IntN(12)
	fmt.Println("randInt", randInt)
	var floatRnd float64
	for i := 0; i <= randInt; i++ {
		floatRnd += rand.Float64()
	}
	fmt.Println("floatRnd at end", floatRnd)
	return floatRnd
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}
