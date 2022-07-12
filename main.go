//Package main helps to learn comment through execises.
package main

import (
	"fmt"
	"myExecises/execises1"
	"myExecises/execises2"
	"myExecises/execises3"
	"myExecises/execises4"
	"myExecises/execises6"
	"myExecises/execises7"
)

func main() {
	fmt.Println("'''''''''''''''''EXE 1 STARTS'''''''''''''''''''''''''''''''''''''''")
	var RemainingTime = execises1.RemainingOvenTime(30)

	fmt.Printf("the remaining time is %v mins \n", RemainingTime)

	prepTime := execises1.PreparationTime(2)
	fmt.Printf("The preparation time is %v mins \n", prepTime)

	elapsTime := execises1.ElapsedTime(3, 20)
	fmt.Printf("The elaps time is %v mins\n", elapsTime)

	fmt.Println("'''''''''''''''''''''EXE 1  ENDS'''''''''''''''''''''''''''''''''''")

	fmt.Println("'''''''''''''''''EXE 2 STARTS'''''''''''''''''''''''''''''''''''''''")

	/*var KnightIsAwake = true
	fmt.Println(execises2.CanFastAttack(KnightIsAwake))
	*/

	/*
			var KnightIsAwake = false
			var archerIsAwake = true
			var prisonerIsAwake = false


		fmt.Println(execises2.CanSpy(KnightIsAwake, prisonerIsAwake, archerIsAwake))
	*/
	var archerIsAwake = false
	var prisonerIsAwake = true

	fmt.Println(execises2.CanSignalPrisoner(archerIsAwake, prisonerIsAwake))

	/*
		var KnightIsAwake = false
		var ArcherIsAwake = true
		var PrisonerIsAwake = false
		var PetDogIsPresent = false
		fmt.Println(execises2.CanFreePresoner(KnightIsAwake, ArcherIsAwake, PrisonerIsAwake, PetDogIsPresent))

	*/
	fmt.Println("'''''''''''''''''''''EXE2  ENDS'''''''''''''''''''''''''''''''''''")

	fmt.Println("''''''''''''''''''''''EXE 3 START''''''''''''''''''''''''''''''''''''''''''''''")
	var greetings = execises3.Welcome("john")

	fmt.Printf("welcome to my party %v\n", greetings)

	var name, age = execises3.HappyBirthday("upright", 45)
	fmt.Printf("happy birthday %v! you are now %v years old \n", name, age)

	var yourName, tableNumber, seatMateName, directon, distanceToTable = execises3.AssignTable("Christain", 27, "frank", "on the left", 23.7834298)

	fmt.Printf("Welcome to my party %v \n", yourName)
	fmt.Printf("you have been assigned to table %03d Your table is %v , exertly %.1f meters from there \n", tableNumber, directon, distanceToTable)
	fmt.Printf("Who You will be sitting next to is %v \n", seatMateName)

	fmt.Println("''''''''''''''''''''''EXE 3 ENDS''''''''''''''''''''''''''''''''''''''''''''''")

	fmt.Println("''''''''''''''''''''''''''EXE 4 STARTS'''''''''''''''''''''''''''''''''''''''''''''''''''")
	var myCad = execises4.ParseCard("ace")
	fmt.Println(myCad)

	myTurn := execises4.FirstTurn("ace", "ace", "jack")

	fmt.Println(myTurn)

	fmt.Println("''''''''''''''''''''''''''EXE 4 ENDS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	fmt.Println("''''''''''''''''''''''''''EXE 6 STARTS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	workingCarParHour := execises6.CalculateWorkingCarsPerHour(1547, 90)
	// fmt.Printf("%.1f \n", workingCarParHour)
	fmt.Println(workingCarParHour)

	workingCarParMinutes := execises6.CalculateWorkingCarsPerMinute(1105, 90)
	fmt.Println(workingCarParMinutes)

	fmt.Println("''''''''''''''''''''''''''EXE 6 ENDS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	fmt.Println("''''''''''''''''''''''''''EXE 7 STARTS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}

	preparationTime := execises7.PreparationTime(layers, 3)

	var myPrepTime = execises7.PreparationTime(layers, 0)

	fmt.Println(preparationTime)
	fmt.Println(myPrepTime)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	eachNoodles, eachSauce := execises7.Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})

	fmt.Printf("%v , %.1f\n", eachNoodles, eachSauce)
	fmt.Println("..........................................")
	execises7.AddSecreteIngredient()
	fmt.Println("..........................................")

	quantities := []float64{1.2, 3.6, 10.5}

	myScaleRecipe := execises7.ScaleRecipe(quantities, 4)

	fmt.Println(myScaleRecipe)

	fmt.Println("''''''''''''''''''''''''''EXE 7 ENDS'''''''''''''''''''''''''''''''''''''''''''''''''''")

}
