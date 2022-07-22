package execises10

import "fmt"

func FavouriteCards() {

	AllCards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s []int

	for _, cards := range AllCards {

		if cards == 2 {
			s = append(s, cards)
		}
		if cards == 6 {
			s = append(s, cards)
		}
		if cards == 9 {
			s = append(s, cards)
		}
	}
	fmt.Printf("%v\n", s)
}

func GetItem(Items []int, IndexPosition int) int {

	for index, myItems := range Items {

		if index == IndexPosition {

			return myItems
		}

	}
	return -1
}

func SetItem(myItems []int, Index, newCard int) []int {

	for i := range myItems {

		if i == Index {
			myItems[i] = newCard
		}
		if -Index == i {
			myItems = append(myItems, newCard)

		}
	}

	if Index > len(myItems)-1 {

		myItems = append(myItems, newCard)

	}
	return myItems
}

func PrependItems(slice []int, index, Card int) []int {

	if index == 0 && Card == 0 {
		return slice
	} else {

		slice = append([]int{index, Card}, slice...)
	}

	return slice

}

func RemoveItem(items []int, index int) []int {

	for i := range items {

		if i == index {

			return append(items[:index], items[index+1:]...)

		}
		if index > len(items)-1 {
			return items
		}

	}
	return items
}
