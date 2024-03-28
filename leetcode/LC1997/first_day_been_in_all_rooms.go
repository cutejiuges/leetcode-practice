/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-28 08:35:15
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-28 08:41:52
 * @FilePath: /leetcode-practice/leetcode/LC1997/first_day_been_in_all_rooms.go
 * @Description:
 */

package lc1997

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := int(1e9 + 7)
	dp := make([]int, len(nextVisit))

	dp[0] = 2
	for i, length := 1, len(nextVisit); i < length; i++ {
		next := nextVisit[i]
		dp[i] = dp[i-1] + 2
		if next != 0 {
			dp[i] = (dp[i] - dp[next-1] + mod) % mod
		}
		dp[i] = (dp[i] + dp[i-1]) % mod
	}
	return dp[len(nextVisit)-2]
}
