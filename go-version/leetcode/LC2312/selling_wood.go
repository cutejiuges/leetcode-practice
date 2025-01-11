/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-15 00:08:57
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-15 00:31:01
 * @FilePath: /leetcode-practice/leetcode/LC2312/selling_wood.go
 * @Description: lc2312，是比较经典的动态规划问题
 */
package lc2312

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int64, n+1)
	}
	for _, price := range prices {
		dp[price[0]][price[1]] = int64(price[2])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			v := int64(dp[i][j])
			for k := 1; k <= i/2; k++ {
				v = M(v, dp[k][j]+dp[i-k][j])
			}
			for k := 1; k <= j/2; k++ {
				v = M(v, dp[i][k]+dp[i][j-k])
			}
			dp[i][j] = v
		}
	}
	return dp[m][n]
}

func M(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
