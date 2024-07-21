package LC1186

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/21 下午1:44
 * @FilePath: maximum_sum
 * @Description: 动态规划题目
 */

func maximumSum(arr []int) int {
	length := len(arr)
	dp := make([][2]int, length)
	dp[0][0], dp[0][1] = arr[0], 0
	ans := arr[0]
	for i := 1; i < length; i++ {
		dp[i][0], dp[i][1] = max(dp[i-1][0], 0)+arr[i], max(dp[i-1][0], dp[i-1][1]+arr[i])
		ans = max(ans, max(dp[i][0], dp[i][1]))
	}
	return ans
}

func maximumSum2(arr []int) int {
	dp0, dp1, ans := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		dp0, dp1 = max(dp0, 0)+arr[i], max(dp0, dp1+arr[i])
		ans = max(dp0, dp1)
	}
	return ans
}
