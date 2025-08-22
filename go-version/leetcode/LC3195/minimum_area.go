package LC3195

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/22 下午8:23
 * @FilePath: minimum_area
 * @Description: 直接遍历
 */

func minimumArea(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	minRow, minCol := m, n
	maxRow, maxCol := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				minRow = min(minRow, i)
				minCol = min(minCol, j)
				maxRow = max(maxRow, i)
				maxCol = max(maxCol, j)
			}
		}
	}
	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
