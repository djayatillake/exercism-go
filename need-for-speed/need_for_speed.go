package speed

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery  int
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	var New Car
	New.speed = speed
	New.batteryDrain = batteryDrain
	New.battery = 100
	return New
}


// Track implements a race track.
type Track struct{
	distance int
}


// NewTrack created a new track
func NewTrack(distance int) Track {
	var New Track
	New.distance = distance
	return New
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move but use the leftover battery.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.distance = car.speed
		car.battery = car.battery - car.batteryDrain
		return car
	} else {
		return car
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	range_remaining := (car.battery / car.batteryDrain) * car.speed
	return track.distance <= range_remaining
}
