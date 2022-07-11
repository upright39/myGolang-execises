package execises3

func Welcome(name string) string {
	return name
}

func HappyBirthday(name string, age int) (string, int) {
	return name, age
}

func AssignTable(yourName string, tableNumber int, seatMateName string, directon string, distanceToTable float64) (string, int, string, string, float64) {

	return yourName, tableNumber, seatMateName, directon, distanceToTable
}
