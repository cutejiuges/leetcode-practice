package LC807

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/14 下午3:38
 * @FilePath: max_increase_keeping_skyline
 * @Description:
 */

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := len(grid)

	rowMaxValue, colMaxValue := make([]int, n), make([]int, n)
	for i, row := range grid {
		for j, height := range row {
			rowMaxValue[i] = max(rowMaxValue[i], height)
			colMaxValue[j] = max(colMaxValue[j], height)
		}
	}

	ans := 0
	for i, row := range grid {
		for j, height := range row {
			ans += min(rowMaxValue[i], colMaxValue[j]) - height
		}
	}
	return ans
}
