package execises11

import "fmt"

func Unit() map[string]int {
	unit := map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozin": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
	return unit

}

func NewBill() map[string]int {

	bill := map[string]int{"carrot": 12}

	return bill

}

func AddItem(bill, units map[string]int, itemName, givenUnit string) bool {
	value, exists := units[givenUnit]
	_, prsent := bill[itemName]

	if !exists {

		fmt.Println("does not exist")
		return exists

	}
	if exists {

		bill[itemName] = value

		// fmt.Printf("if exist: %v\n", bill)

	} else if prsent {

		bill[itemName] += value

		// fmt.Printf("if else: %v", bill)

	}

	return true
}

func RemoveItem(bill, units map[string]int, itemName, givenUnit string) bool {
	_, itemExists := bill[itemName]
	value, unitExists := units[givenUnit]

	if !itemExists {

		return itemExists

	}
	if !unitExists {

		return unitExists
	}
	if value < 0 {

		return false
	}
	if value == 0 {

		return true
	} else {

		delete(bill, itemName)

		units[givenUnit]--

	}

	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	value, exist := bill[item]

	if !exist {

		return value, exist
	} else {
		return value, exist
	}
}
