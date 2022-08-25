//Package execises1 contains or provides tools to learn the basics of GO language.
package execises1

//OvenTime a constant defining the given oven time
const OvenTime = 40

//RemainingOvenTime function takes and integer and returns an integer which is the remainder of the oven time.
func RemainingOvenTime(actual int) int {

	return OvenTime - actual
}

//PreparationTime function takes and integer and  returns the PreparationTime.
func PreparationTime(numberOfLayers int) int {
	eachLayerInMinutes := 2
	return eachLayerInMinutes * numberOfLayers
}

//ElapsedTime function takes in two paramiters and return and integer with is the duration of time.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	sumOfPreparatintime := numberOfLayers * 2

	timeSpentInMinutes := 40 - actualMinutesInOven

	sum := sumOfPreparatintime + timeSpentInMinutes

	return sum

}
