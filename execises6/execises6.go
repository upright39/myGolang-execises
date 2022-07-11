//Package execises6 provides tools for calculating time used to produce a car.
package execises6

//CalculateWorkingCarsPerHour function takes in two integer paramiter and return sucessful car produce in hours in integer number.

func CalculateWorkingCarsPerHour(nubmerOfCarsProduseParHour, sucessRate int) float64 {

	succesfulCarsPerHour := (float64(nubmerOfCarsProduseParHour) * float64(sucessRate)) / 100

	return succesfulCarsPerHour

}

//CalculateWorkingCarsPerMinute function takes in two integer parameter and return sucessful car produce in minutes in integer number.

func CalculateWorkingCarsPerMinute(nubmerOfCarsProduseParHour, sucessRate int) int {
	succesfulCarsPerMinute := float64(nubmerOfCarsProduseParHour / 60 * sucessRate / 100)

	return int(succesfulCarsPerMinute)
}
