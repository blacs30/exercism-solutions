package jedlik

import "fmt"

/*
Implement the `Drive` method on the `Car` that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage:
*/
func (c *Car) Drive() {
	if c.battery-c.batteryDrain >= 0 {
		c.distance = c.distance + c.speed
		c.battery = c.battery - c.batteryDrain
	}
}

/*
Implement a `DisplayDistance` method on `Car` to return the distance as displayed on the LED display as a `string`:
*/
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

/*
To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line.
Implement the `CanFinish` method that takes a `trackDistance int` as its parameter and returns `true` if the car can finish the race; otherwise, return `false`:
*/
func (c Car) CanFinish(trackDistance int) bool {
	return trackDistance <= c.battery/c.batteryDrain*c.speed
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
