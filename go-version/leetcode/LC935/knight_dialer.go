package LC935

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/10 下午11:10
 * @FilePath: knight_dialer
 * @Description: 动态规划
 */

func knightDialer(n int) int {
	mod := 1000000007
	nextPos := [][]int{{4, 6}, {8, 6}, {7, 9}, {4, 8}, {9, 0, 3}, {}, {0, 7, 1}, {6, 2}, {1, 3}, {2, 4}}
	dp := make([]int, 10)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		cur := make([]int, 10)
		for j := 0; j < 10; j++ {
			for _, next := range nextPos[j] {
				cur[j] = (cur[j] + dp[next]) % mod
			}
		}
		dp = cur
	}
	ans := 0
	for i := range dp {
		ans = (ans + dp[i]) % mod
	}
	return ans
}
