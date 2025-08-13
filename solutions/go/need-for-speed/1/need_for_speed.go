package speed

type Car struct {
	speed, batteryDrain, battery, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		battery:      100,
		batteryDrain: batteryDrain,
		distance:     0,
	}

}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	} else {
		return car
	}
	return car

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	drivenCar := car
	for drivenCar.battery >= drivenCar.batteryDrain {
		drivenCar = Drive(drivenCar)
	}
	return drivenCar.distance >= track.distance
}
