/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-25 08:59:07
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-25 09:02:47
 * @FilePath: /leetcode-practice/leetcode/LC518/change.go
 * @Description: 零钱兑换组合数问题，直接叠加即可
 */
package lc518

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
