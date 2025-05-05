package LC790

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/5 下午9:02
 * @FilePath: num_tilings
 * @Description:
 */

func numTilings(n int) int {
	mod := int(1e9) + 7
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[1][0], dp[1][1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][1]
		var cur int
		for j := 0; j < 4; j++ {
			cur = (cur + dp[i-1][j]) % mod
		}
		dp[i][1] = cur
		dp[i][2] = (dp[i-1][0] + dp[i-1][3]) % mod
		dp[i][3] = (dp[i-1][0] + dp[i-1][2]) % mod
	}
	return dp[n][1]
}
