package LC1547

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/11 下午11:08
 * @FilePath: min_cost
 * @Description:
 */

// 区间dp
func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)

	m := len(cuts)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := m - 3; i >= 0; i-- {
		for j := i + 2; j < m; j++ {
			res := 0x3f3f3f3f
			for k := i + 1; k < j; k++ {
				res = min(res, dp[i][k]+dp[k][j])
			}
			dp[i][j] = res + cuts[j] - cuts[i]
		}
	}
	return dp[0][m-1]
}
