/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-24 11:51:06
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-24 11:58:11
 * @FilePath: /leetcode-practice/leetcode/LC322/coin_change.go
 * @Description: lc日题练习，零钱兑换，经典的动态规划背包问题
 */
package lc322

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	length := len(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < length; j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
