package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	/*
		| Unit               | Score |
		| ------------------ | ----- |
		| quarter_of_a_dozen | 3     |
		| half_of_a_dozen    | 6     |
		| dozen              | 12    |
		| small_gross        | 120   |
		| gross              | 144   |
		| great_gross        | 1728  |
	*/
	uom := map[string]int{}
	uom["quarter_of_a_dozen"] = 3
	uom["half_of_a_dozen"] = 6
	uom["dozen"] = 12
	uom["small_gross"] = 120
	uom["gross"] = 144
	uom["great_gross"] = 1728
	return uom
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// Check if the unit exists in the units map
	// assign the value to be added to the bill
	value, exists := units[unit]
	if !exists {
		return false
	}
	itemValue, exists := bill[item]
	if exists {
		bill[item] = itemValue + value

	} else {
		bill[item] = value
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	/*
	   Task:
	   - Return `false` if the given item is **not** in the bill
	   - Return `false` if the given `unit` is not in the `units` map.
	   - Return `false` if the new quantity would be less than 0.
	   - If the new quantity is 0, completely remove the item from the `bill` then return `true`.
	   - Otherwise, reduce the quantity of the item and return `true`.
	*/

	itemValue, exists := bill[item]
	if !exists {
		return false
	}
	unitValue, exists := units[unit]
	if !exists {
		return false
	}
	if itemValue < unitValue {
		return false
	} else if itemValue == unitValue {
		delete(bill, item)
		return true
	} else {
		bill[item] = itemValue - unitValue
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	/*
	   Task:

	   - Return `0` and `false` if the `item` is not in the bill.
	   - Otherwise, return the quantity of the item in the `bill` and `true`.
	*/

	value, exists := bill[item]
	if !exists {
		return 0, false
	} else {
		return value, true
	}

}
