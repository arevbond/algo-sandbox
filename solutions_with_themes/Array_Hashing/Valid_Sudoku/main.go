package main

var allowedChars []byte = []byte{'1', '2', '3', '4', '5', '6',
	'7', '8', '9', '.'}

func checkRow(row []byte) bool {
	seen := make(map[byte]bool)
	for _, ch := range row {
		ok1 := checkIn(ch, allowedChars)
		_, ok2 := seen[ch]
		if !ok1 || ok2 {
			return false
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for _, row := range board {
		if !checkRow(row) {
			return false
		}
	}

	for i := 0; i < len(board); i++ {
		col := make([]byte, 0)
		for j := 0; j < 9; j++ {
			col = append(col, board[i][j])
		}
		if !checkRow(col) {
			return false
		}
	}
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			row := make([]byte, 0)
			for k := i; k < k+3; k++ {
				for t := j; t < t+3; t++ {
					row = append(row, board[k][t])
				}
			}
			if !checkRow(row) {
				return false
			}
		}
	}
	return true
}

func checkIn(x byte, slice []byte) bool {
	for _, ch := range slice {
		if ch == x {
			return true
		}
	}
	return false
}
