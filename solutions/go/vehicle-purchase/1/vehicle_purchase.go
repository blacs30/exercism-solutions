package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return fmt.Sprintf("%s is clearly the better choice.", option1)
	} else {
		return fmt.Sprintf("%s is clearly the better choice.", option2)

	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= float64(10) {
		return originalPrice * float64(50) / float64(100)
	} else if age >= float64(3) {
		return originalPrice * float64(70) / float64(100)
	} else {
		return originalPrice * float64(80) / float64(100)
	}
}
