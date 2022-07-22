package execises12

// func CountInRank(board map[string]int, rank string) int {
// 	value, ok := board[rank]

// 	if ok {
// 		return value
// 	}
// 	return 0
// }

// else if !ok {
// 	return value

// }

func CountInRank(board map[string]int, rank string) int {
	value, exist := board[rank]

	for range board {

		if exist {
			return value
		}

	}
	return 0
}

func CounterInFile(board map[int]int, file int) int {

	value, exist := board[file]

	for {

		if exist {

			return value
		} else if !exist {
			return value
		}
	}
}

func CountAll(board map[string]int) int {
	count := 0
	for i := range board {

		if len(board) > 0 {

			count += board[i] * 2
		} else {
			return 0
		}

	}
	return count
}

func CountOccupied(board map[string]int) int {
	count := 0
	for i := range board {

		if len(board) > 0 {

			count += board[i]
		}

	}
	return count
}
