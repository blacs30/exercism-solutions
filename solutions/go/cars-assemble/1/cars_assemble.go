package cars

//import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	succ := float64(productionRate) * successRate / 100
	return succ
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	hourlyCars := int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
	return hourlyCars
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupBuildCost := uint(carsCount) / uint(10) * uint(95000)
	singleBuildCost := uint(carsCount) % uint(10) * uint(10000)
	return groupBuildCost + singleBuildCost

}
