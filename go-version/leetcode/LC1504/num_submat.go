package LC1504

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/21 下午10:54
 * @FilePath: num_submat
 * @Description:
 */

func numSubmat(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	m, n := len(mat), len(mat[0])
	temp := make([][]int, m)
	for i := range temp {
		temp[i] = make([]int, n)
	}
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				temp[i][j] = mat[i][j]
			} else {
				if mat[i][j] == 0 {
					temp[i][j] = 0
				} else {
					temp[i][j] = temp[i][j-1] + 1
				}
			}
			cur := temp[i][j]
			for k := i; k >= 0; k-- {
				cur = min(cur, temp[k][j])
				if cur == 0 {
					break
				}
				ans += cur
			}
		}
	}
	return ans
}
