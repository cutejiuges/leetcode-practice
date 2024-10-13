package LC1884

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/13 下午4:47
 * @FilePath: two_egg_drop
 * @Description: 经典的面试题，鸡蛋掉落的问题
 */

func twoEggDrop(n int) int {
	//使用动态规划的思路
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 0x3f3f3f
	}
	//第0层一定不会碎，用0个鸡蛋即可
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			dp[i] = min(dp[i], max(k-1, dp[i-k])+1)
		}
	}
	return dp[n]
}
