package arrayshashing

func isValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	sqrs := make([]int, 9)

	for ri := 0; ri < 9; ri++ {
		for ci := 0; ci < 9; ci++ {
			if board[ri][ci] == '.' {
				continue
			}

			val := board[ri][ci] - '1'
			bit := 1 << val
			sqrsi := (ri/3)*3 + ci/3

			if rows[ri]&bit != 0 || cols[ci]&bit != 0 ||
				sqrs[sqrsi]&bit != 0 {
				return false
			}

			rows[ri] |= bit
			cols[ci] |= bit
			sqrs[sqrsi] |= bit
		}
	}

	return true
}
