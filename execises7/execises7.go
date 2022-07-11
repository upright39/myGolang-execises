package execises7

import "fmt"

func PreparationTime(layers []string, avePrepTimeInMins int) int {

	if avePrepTimeInMins == 0 {
		avePrepTimeInMins = 2
		prepTime := len(layers) * avePrepTimeInMins
		return prepTime
	} else {

		prepTime := len(layers) * avePrepTimeInMins
		return prepTime
	}
}

func Quantities(myLasanga []string) (int, float64) {

	eachNoodles := 0
	eachSauce := 0.0
	var initialNoodleContaner []string
	var initialSauceContaner []string

	// for i := 0; i < len(myLasanga); i++ {

	// 	if myLasanga[i] == "noodles" {
	// 		nuberOfNoodlesInLasanga := append(initialNoodleContaner, myLasanga[i])

	// 		eachNoodles += len(nuberOfNoodlesInLasanga) * 50
	// 	}
	// 	if myLasanga[i] == "sauce" {

	// 		nuberOfSauceInLasanga := append(initialSauceContaner, myLasanga[i])

	// 		eachSauce += float64(len(nuberOfSauceInLasanga)) * 0.2

	// 	}
	// }
	// return eachNoodles, eachSauce

	for _, Lasanga := range myLasanga {

		if Lasanga == "noodles" {
			nuberOfNoodlesInLasanga := append(initialNoodleContaner, Lasanga)
			eachNoodles += len(nuberOfNoodlesInLasanga) * 50
		}

		if Lasanga == "sauce" {

			nuberOfSauceInLasanga := append(initialSauceContaner, Lasanga)

			eachSauce += float64(len(nuberOfSauceInLasanga)) * 0.2

		}

	}
	return eachNoodles, eachSauce
}

func AddSecreteIngredient() {
	friendList := []string{"noodles", "sauce", "mozzarella", "kampot papper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	if len(friendList) > 0 {

		friendLists := friendList[len(friendList)-1]

		myList = myList[:len(myList)-1]

		myList = append(myList, friendLists)

	}

	fmt.Println(myList)

}

func ScaleRecipe(Quantities []float64, numPortions int) []float64 {

	for _, Quant := range Quantities {

		Quant

	}
	return Quant
}
