package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidRow(board[i]) {
			return false
		}

		vRow := []byte{board[0][i], board[1][i], board[2][i], board[3][i], board[4][i], board[5][i],
			board[6][i], board[7][i], board[8][i]}

		if !isValidRow(vRow) {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boxRow := []byte{
				board[0+i*3][0+j*3], board[0+i*3][1+j*3], board[0+i*3][2+j*3],
				board[1+i*3][0+j*3], board[1+i*3][1+j*3], board[1+i*3][2+j*3],
				board[2+i*3][0+j*3], board[2+i*3][1+j*3], board[2+i*3][2+j*3],
			}

			if !isValidRow(boxRow) {
				return false
			}
		}
	}

	return true
}

func isValidByte(b byte) bool {
	if b == '.' || b == '1' || b == '2' || b == '3' || b == '4' || b == '5' || b == '6' ||
		b == '7' || b == '8' || b == '9' {
		return true
	}
	return false
}

func isValidRow(row []byte) bool {

	for i := 0; i < 9; i++ {
		if row[i] == '.' {
			continue
		}

		if !isValidByte(row[i]) {
			return false
		}

		if i < 9 {
			for k := i + 1; k < 9; k++ {
				if row[k] != '.' && row[i] == row[k] {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
		{'2', '3', '.', '.', '.', '.', '.', '.', '.'},
		{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	result := isValidSudoku(board)
	println(result)
}
