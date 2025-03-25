package LC2711

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/25 下午11:03
 * @FilePath: difference_of_distinct_values
 * @Description: 遍历数组模拟计算
 */

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bottomRightSet := make(map[int]struct{})
			x, y := i+1, j+1
			for x < m && y < n {
				bottomRightSet[grid[x][y]] = struct{}{}
				x++
				y++
			}
			topLeftSet := make(map[int]struct{})
			x, y = i-1, j-1
			for x >= 0 && y >= 0 {
				topLeftSet[grid[x][y]] = struct{}{}
				x--
				y--
			}
			ans[i][j] = abs(len(bottomRightSet) - len(topLeftSet))
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
