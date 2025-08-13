package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
/*
When practicing with her cards, Elyse likes to start with her favorite three cards of the deck: 2, 6 and 9.
Write a function `FavoriteCards` that returns a slice with those cards in that order.
*/
func FavoriteCards() []int {
	return []int{2, 6, 9}

}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
/*
Return the card at position `index` from the given stack.
*/
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
/*
Exchange the card at position `index` with the new card provided and return the adjusted stack.
Note that this will modify the input slice which is the expected behavior.
*/
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
/*
Add the card(s) specified in the `value` parameter at the top of the stack.

```go
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice, 5, 1)
fmt.Println(cards)
// Output: [5 1 3 2 6 4 8]
*/
func PrependItems(slice []int, values ...int) []int {
	var newSlice []int
	newSlice = values
	newSlice = append(newSlice, slice...)
	return newSlice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
/*

Remove the card at position `index` from the stack and return the stack.
Note that this may modify the input slice which is ok.

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
fmt.Println(cards)
// Output: [3 2 4 8]
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:
*/
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	} else {
		newSlice := slice[:index]
		newSlice = append(newSlice, slice[index+1:]...)
		return newSlice
	}
}
