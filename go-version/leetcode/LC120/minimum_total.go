package LC120

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/25 下午11:54
 * @FilePath: minimum_total
 * @Description: 经典动态规划，杨辉三角
 */

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := dp[n-1][0]
	for i := 1; i < n; i++ {
		ans = min(ans, dp[n-1][i])
	}
	return ans
}
