package LC1277

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/20 上午8:39
 * @FilePath: count_square
 * @Description: 动态规划
 */

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			ans += dp[i][j]
		}
	}
	return ans
}
