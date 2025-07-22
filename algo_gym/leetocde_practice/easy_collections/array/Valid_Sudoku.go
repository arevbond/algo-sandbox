func isValidSudoku(board [][]byte) bool {
	for _, row := range board {
		if !isCorrectLine(row) {
			return false
		}
	}
	 
	for col := 0; col < len(board[0]); col++ {
		line := make([]byte, 0, len(board))
		for row := 0; row < len(board); row++ {
			line = append(line, board[row][col])
		}
		if !isCorrectLine(line) {
			return false
		}
	}

	for startRow := 0; startRow < len(board); startRow += 3 {
		startCol := 0
		line := make([]byte, 0, 9)
		for row := startRow; row < startRow + 3 && row < len(board); row++ {
			for col := startCol; col < startCol + 3 && col < len(board[0]); col++ {
				line = append(line, board[row][col])
			}

			if !isCorrectLine(line) {
				return false
			}
		}
	}

	for startCol := 0; startCol < len(board); startCol += 3 {
		startRow := 0
		line := make([]byte, 0, 9)
		for row := startRow; row < startRow + 3 && row < len(board); row++ {
			for col := startCol; col < startCol + 3 && col < len(board[0]); col++ {
				line = append(line, board[row][col])
			}

			if !isCorrectLine(line) {
				return false
			}
		}
	}
	
	for startCol := 3; startCol < len(board); startCol += 3 {
		for startRow := 3; startRow < len(board); startRow += 3 {
			line := make([]byte, 0, 9)
			for row := startRow; row < startRow + 3 && row < len(board); row++ {
				for col := startCol; col < startCol + 3 && col < len(board[0]); col++ {
					line = append(line, board[row][col])
				}

				if !isCorrectLine(line) {
					return false
				}
			}
		}
	}

	return true

}

func isCorrectLine(line []byte) bool {
	set := make(map[byte]struct{})

	for _, val := range line {
		if val == '.' {
			continue 
		}

		if _, ok := set[val]; ok {
			return false
		}
		set[val] = struct{}{}
	}

	return true
}
