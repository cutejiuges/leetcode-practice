package lc1706

func findBall(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		curCol := i
		for _, row := range grid {
			dir, nextCol := row[i], curCol+row[i]
			if nextCol < 0 || nextCol >= n || row[nextCol] != dir {
				curCol = -1
				break
			}
			curCol = nextCol
		}
		ans[i] = curCol
	}
	return ans
}
