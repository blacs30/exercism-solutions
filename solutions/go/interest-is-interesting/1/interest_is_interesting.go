package interest

/*
- 3.213% for a balance less than `0` dollars (balance gets more negative).
- 0.5% for a balance greater than or equal to `0` dollars, and less than `1000` dollars.
- 1.621% for a balance greater than or equal to `1000` dollars, and less than `5000` dollars.
- 2.475% for a balance greater than or equal to `5000` dollars.
Implement the `InterestRate()` function to calculate the interest rate based on the specified balance:
*/
// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < float64(0):
		return float32(3.213)
	case balance >= float64(5000):
		return float32(2.475)
	case balance >= float64(1000):
		return float32(1.621)
	default:
		return float32(0.5)
	}
}

// Implement the `Interest()` function to calculate the interest based on the specified balance:
// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + (balance * float64(InterestRate(balance)) / 100)
}

/*
Implement the `YearsBeforeDesiredBalance()` function to calculate the minimum number of years required to reach the desired balance, taking into account that each year, interest is added to the balance.
This means that the balance after one year is: start balance + interest for start balance.
The balance after the second year is: balance after one year + interest for balance after one year.
And so on, until the current year's balance is greater than or equal to the target balance.
*/
// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yearsCounter := 0

	for {
		if balance >= targetBalance {
			break
		}
		balance += Interest(balance)
		yearsCounter++
	}

	return yearsCounter
}
