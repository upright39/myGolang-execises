//Package main helps to learn comment through execises.
package main

import (
	"fmt"
	"myExecises/execises1"
	"myExecises/execises10"
	"myExecises/execises11"
	"myExecises/execises12"
	"myExecises/execises13"
	"myExecises/execises14"
	"myExecises/execises15"
	"myExecises/execises16"
	"myExecises/execises17"
	"myExecises/execises2"
	"myExecises/execises3"
	"myExecises/execises4"
	"myExecises/execises6"
	"myExecises/execises7"
	"myExecises/execises8"
	"myExecises/execises9"
	"time"
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

	fmt.Println("''''''''''''''''''''''''''EXE 8 STARTS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	var myString = "🏃recommended search product 👎"

	fmt.Println(execises8.Application(myString))

	fmt.Println(".............................................................")

	log := "please replace '👎' with '👍'"

	myReplace := execises8.Replace(log, '👎', '👍')

	fmt.Printf("please replace '%c' with '%c' \n", myReplace, myReplace)

	fmt.Println(".............................................................")

	IsWithinLimit := execises8.WithinLimit("hello🏃", 6)

	fmt.Printf("%v\n", IsWithinLimit)

	fmt.Println("''''''''''''''''''''''''''EXE 8 ENDS'''''''''''''''''''''''''''''''''''''''''''''''''''")
	fmt.Println("''''''''''''''''''''''''''EXE 9 STARTS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	execises9.Schedule("07/20/1990 13:45:00")

	fmt.Println("..............................................................")

	execises9.HasPassed("July 25, 2019 13:45:00")
	fmt.Println("..............................................................")
	execises9.IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
	fmt.Println("..............................................................")

	execises9.Description("07/25/2019 13:45:00")

	fmt.Println("..............................................................")
	MyAniversary := execises9.AniversaryDate()
	fmt.Println(MyAniversary)
	fmt.Println("''''''''''''''''''''''''''EXE 9 ENDS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	fmt.Println("''''''''''''''''''''''''''EXE 10 STARTS'''''''''''''''''''''''''''''''''''''''''''''''''''")
	execises10.FavouriteCards()
	fmt.Println("..............................................................")

	myItems := execises10.GetItem([]int{1, 2, 4, 1}, 3)
	fmt.Println(myItems)

	fmt.Println("..............................................................")

	myNewItem := execises10.SetItem([]int{1, 2, 4, 1}, 7, 6)

	fmt.Println(myNewItem)
	fmt.Println("..............................................................")

	Prepend := execises10.PrependItems([]int{3, 2, 6, 4, 8}, 3, 4)
	fmt.Println(Prepend)
	fmt.Println("..............................................................")

	removeItem := execises10.RemoveItem([]int{3, 2, 6, 4, 8}, 11)
	fmt.Println(removeItem)
	fmt.Println("''''''''''''''''''''''''''EXE 10 ENDS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	fmt.Println("''''''''''''''''''''''''''EXE 11 STARTS'''''''''''''''''''''''''''''''''''''''''''''''''''")

	units := execises11.Unit()
	fmt.Println(units)

	fmt.Println("..............................................................")

	bills := execises11.NewBill()
	fmt.Println(bills)

	fmt.Println("..............................................................")

	myBil := execises11.NewBill()
	myUnit := execises11.Unit()

	items := execises11.AddItem(myBil, myUnit, "carrot", "dozen")

	fmt.Println(items)

	fmt.Println("..............................................................")
	myBl := execises11.NewBill()
	myUt := execises11.Unit()

	remove := execises11.RemoveItem(myBl, myUt, "carrot", "dozen")

	fmt.Println(remove)

	fmt.Println("..............................................................")

	bill := map[string]int{"carrot": 12, "grapes": 3}

	qty, ok := execises11.GetItem(bill, "carrot")

	fmt.Println("output:", qty)

	fmt.Println("output:", ok)

	fmt.Println(".....................'EXE 12 STARTS'.........................................")

	chessedBoard := map[string]int{"A": 6, "B": 6, "C": 5, "D": 4}

	countRank := execises12.CountInRank(chessedBoard, "A")

	fmt.Println(countRank)

	fmt.Println("..............................................................")

	Board := map[int]int{2: 5, 3: 4, 4: 4, 5: 3, 6: 2, 7: 1, 8: 0}

	countinfile := execises12.CounterInFile(Board, 2)

	fmt.Println(countinfile)

	fmt.Println("..............................................................")

	chessed := map[string]int{"A": 6, "B": 6, "C": 5, "D": 4, "E": 3, "F": 2, "G": 1, "H": 0}

	count := execises12.CountAll(chessed)

	fmt.Println(count)

	fmt.Println("..............................................................")

	chessedB := map[string]int{"A": 3, "B": 1, "C": 1, "D": 0, "E": 2, "F": 0, "G": 1, "H": 7}

	CountOccupied := execises12.CountOccupied(chessedB)

	fmt.Println(CountOccupied)

	fmt.Println("......................EXE 13 STARTS........................................")

	speed := 5
	batteryDrain := 2

	new_car := execises13.NewCar(speed, batteryDrain)

	fmt.Printf("NewCars:%v\n", new_car)

	fmt.Println("..............................................................")

	destance := 800
	Track := execises13.NewTrack(destance)

	fmt.Printf("Track Destance: %v\n", Track)

	fmt.Println("..............................................................")
	newCars := execises13.NewCar(speed, batteryDrain)
	drive := execises13.Drive(newCars)
	fmt.Println(drive)

	fmt.Println("...................................................................................")
	speeds := 5
	batteryDrains := 2
	nCar := execises13.NewCar(speeds, batteryDrains)

	dest := 100
	tracks := execises13.NewTrack(dest)

	CarFinish := execises13.CarFinish(nCar, tracks)

	fmt.Println(CarFinish)
	fmt.Println("......................EXE 13 ENDS........................................")
	fmt.Println("......................EXE 14 START HERE........................................")

	spds := 5
	batteryD := 2

	car := execises14.NewCar(spds, batteryD)

	fmt.Printf("car is now car: %v\n", car.Drive())

	fmt.Println("'''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	dist := car.DesplayDistance()
	fmt.Println(dist)

	fmt.Println("'''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	desplayBattery := car.DesplayBattery()

	fmt.Println(desplayBattery)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")
	trackDis := 100
	carFinish := car.CanFinish(trackDis)
	fmt.Println(carFinish)

	fmt.Println("......................EXE 14 ENDS HERE........................................")

	fmt.Println("......................EXE 15 STARTS HERE........................................")

	initialVote := 2

	// var counter *int
	var counter = execises15.NewVoteCounter(initialVote)

	fmt.Println(*counter == initialVote)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	votes := 3

	// var VoteCounter *int
	var VoteCounter = &votes

	fmt.Println(execises15.VoteCount(VoteCounter))

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	execises15.IncrementVoteCount(VoteCounter, 2)
	fmt.Println(votes == 5)
	fmt.Println(*VoteCounter == 5)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	result := execises15.NewElectionResult("peter", 3)
	fmt.Println(result.Name == "peter")
	fmt.Println(result.Votes == 3)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	var results = &execises15.ElectionResult{Name: "john", Votes: 32}

	desplayResult := execises15.DesplayResult(results)
	fmt.Println(desplayResult)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	var finalResult = map[string]int{"john": 51, "mary": 10}

	execises15.DecrementVotesOfCandidate(finalResult, "mary")

	fmt.Println(finalResult["mary"])

	fmt.Println("......................EXE 16 STARTS........................................")
	names := "mattew sanabria"
	ages := 29
	addresss := map[string]string{"street": "main st."}

	newResident := execises16.NewResident(names, ages, addresss)
	fmt.Println(newResident)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	my_names := "mattew sanabria"
	my_ages := 4
	my_addresss := make(map[string]string)

	//    or
	//my_addresss := map[string]string{}

	newRes := execises16.NewResident(my_names, my_ages, my_addresss)

	fmt.Println(newRes.HasRequiredInfo())

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	my_name := "mattew sanabria"
	my_age := 4
	my_addres := map[string]string{"street": "16 eberi omuma"}

	newResi := execises16.NewResident(my_name, my_age, my_addres)

	newResi.Delete()

	fmt.Println(newResi)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	name1 := "upright man"
	age1 := 45
	address1 := map[string]string{"street": "no 16 aba road"}

	resident1 := execises16.NewResident(name1, age1, address1)

	name2 := "upright man"
	age2 := 45
	address2 := map[string]string{}

	resident2 := execises16.NewResident(name2, age2, address2)

	residents := []*execises16.Resident{resident1, resident2}

	countNum := execises16.Count(residents)
	fmt.Println(countNum)
	fmt.Println("'''''''''''''''''''''''''''''''''EXE 17 ON CHANNEL'''''''''''''''''''''''''''''''''''''''''''''''''")

	done := make(chan bool)
	go execises17.Hello(done)
	<-done
	fmt.Println("main function")

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	mydone := make(chan bool)

	fmt.Println("the main is going to call the chantuto go routine")
	go execises17.MyHello(mydone)

	<-mydone
	fmt.Println("main receives data ")

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	number := 123
	sqrch := make(chan int)

	cubech := make(chan int)

	go execises17.CalcSquares(number, sqrch)

	go execises17.CalcCubes(number, cubech)

	squares, cubes := <-sqrch, <-cubech

	fmt.Println("Final output", squares+cubes)

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	myCounts := make(chan int)

	go execises17.Mycount(myCounts)

	for v := range myCounts {

		fmt.Println("Receive", v)
	}

	fmt.Println("''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''")

	ch := make(chan int, 2)
	go execises17.Write(ch)
	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
