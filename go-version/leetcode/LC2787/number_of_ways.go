package LC2787

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/12 下午11:05
 * @FilePath: number_of_ways
 * @Description: 01背包+快速幂
 */

func numberOfWays(n, x int) int {
	mod := int(1e9 + 7)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; quickPower(i, x) <= n; i++ {
		val := quickPower(i, x)
		for j := n; j >= val; j-- {
			dp[j] = (dp[j] + dp[j-val]) % mod
		}
	}
	return dp[n]
}

func quickPower(a, b int) int {
	ans := 1
	for b > 0 {
		if b&1 == 1 {
			ans *= a
		}
		a *= a
		b >>= 1
	}
	return ans
}
