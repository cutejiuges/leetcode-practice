package lc59

var dir = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	curNum, maxNum := 1, n*n
	curDir := 0
	curRow, curCol := 0, 0
	for ; curNum <= maxNum; curNum++ {
		matrix[curRow][curCol] = curNum
		if r, c := curRow+dir[curDir][0], curCol+dir[curDir][1]; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] != 0 {
			curDir = (curDir + 1) % 4
		}
		curRow += dir[curDir][0]
		curCol += dir[curDir][1]
	}
	return matrix
}
