package LC2732

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/25 下午10:35
 * @FilePath: good_subset_of_binary_matrix
 * @Description:
 */

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	ans := make([]int, 0)
	mp := make(map[int]int)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		st := 0
		for j := 0; j < n; j++ {
			st |= grid[i][j] << j
		}
		mp[st] = i
	}

	if _, ok := mp[0]; ok {
		ans = append(ans, mp[0])
		return ans
	}

	for x, i := range mp {
		for y, j := range mp {
			if x&y == 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return ans
}
