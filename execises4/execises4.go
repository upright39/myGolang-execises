package execises4

func ParseCard(myCard string) int {

	switch myCard {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "other":
		return 0
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {

	switch {
	case card1 == "ace" && card2 == "ace", dealerCard == "jack":
		return "p"

	case card1 == "ace" && card2 == "king", dealerCard != "ace", dealerCard != "jack", dealerCard != "queen", dealerCard != "ten":
		return "s"
		//how will i do range in a data type of sring
	case card1 == "five" && card2 == "quen", dealerCard == "ace":
		return "H"
	default:
		return "imputs value mismatchd"
	}

}
