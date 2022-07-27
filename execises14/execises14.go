package execises14

import "fmt"

type Car struct {
	speed, batteryDrain, battery, distance int
}

func NewCar(speeds, batteryDrains int) Car {

	car := Car{
		speed:        speeds,
		batteryDrain: batteryDrains,
		battery:      100,
	}
	return car
}

func (c Car) Drive() Car {

	d := c.battery - c.batteryDrain
	c.battery = d
	c.distance += c.speed

	return c
}

func (c Car) DesplayDistance() string {

	return fmt.Sprintf("output: Driven %v meters", c.distance)
}

func (c Car) DesplayBattery() string {
	return fmt.Sprint("output: Battery at ", c.battery, "%")
}

func (c Car) CanFinish(trackDistance int) bool {
	if c.speed < trackDistance {

		return true
	} else {
		return false
	}
}
