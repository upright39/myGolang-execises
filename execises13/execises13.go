package execises13

type Car struct {
	speed, batteryDrain, battery, distance int
}

func NewCar(speeds, batteryDrains int) Car {

	newCar := Car{
		speed:        speeds,
		batteryDrain: batteryDrains,
		battery:      100,
	}
	return newCar
}

type Track struct {
	distance int
}

func NewTrack(destance int) Track {

	Dis := Track{
		distance: destance,
	}
	return Dis
}

func Drive(NewCar Car) Car {

	d := NewCar.battery - NewCar.batteryDrain
	NewCar.battery = d

	NewCar.speed += NewCar.distance

	NewCar.distance += NewCar.speed

	return NewCar
}

func CarFinish(NewCar Car, NewTrack Track) bool {

	if NewCar.speed < NewTrack.distance {
		return true
	} else {
		return false
	}
}
