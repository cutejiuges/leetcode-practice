package LC887

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/14 下午11:04
 * @FilePath: super_egg_drop
 * @Description:
 */

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	//枚举扔鸡蛋的次数
	m := 0
	for dp[K][m] < N {
		m++
		for i := 1; i <= K; i++ {
			dp[i][m] = dp[i-1][m-1] + dp[i][m-1] + 1
		}
	}
	return m
}
