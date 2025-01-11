package LC2713

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/19 下午10:57
 * @FilePath: max_increasing_cells
 * @Description: 动态规划咯
 */
func maxIncreasingCells(mat [][]int) int {
	row, col := len(mat), len(mat[0])
	mp := make(map[int][][2]int)
	rows, cols := make([]int, row), make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			mp[mat[i][j]] = append(mp[mat[i][j]], [2]int{i, j})
		}
	}
	keys := make([]int, 0)
	for key, _ := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		curPos := mp[key]
		sameField := make([]int, 0)
		for _, arr := range curPos {
			sameField = append(sameField, max(rows[arr[0]], cols[arr[1]])+1)
		}

		for i := 0; i < len(curPos); i++ {
			arr, d := curPos[i], sameField[i]
			rows[arr[0]] = max(rows[arr[0]], d)
			cols[arr[1]] = max(cols[arr[1]], d)
		}
	}

	return maxValue(rows)
}

func maxValue(array []int) int {
	ans := array[0]
	for _, v := range array {
		if v > ans {
			ans = v
		}
	}
	return ans
}
