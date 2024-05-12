package LC1553

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/12 下午7:49
 * @FilePath: min_days
 * @Description: 经典动态规划
 */

// 直接动规，空间超限，时间超限
func minDays(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
	}
	return dp[n]
}

// 取余优化一版，使用记忆化搜索，通过
func minDays2(n int) int {
	dp := make(map[int]int)
	dp[0], dp[1] = 0, 1

	var search func(int) int

	search = func(n int) int {
		if day, ok := dp[n]; ok {
			return day
		}
		dp[n] = min(n%2+1+search(n/2), n%3+1+search(n/3))
		return dp[n]
	}

	return search(n)
}
