package LC2140

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/1 下午10:48
 * @FilePath: most_points
 * @Description:
 */

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], int64(questions[i][0])+dp[min(n, i+questions[i][1]+1)])
	}
	return dp[0]
}
