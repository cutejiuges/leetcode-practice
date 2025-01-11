package LC1463

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/7 上午12:11
 * @FilePath: cherry_pickup
 * @Description: 多维动态规划
 */

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dp[0][n-1] = grid[0][0] + grid[0][n-1]

	nxt := make([][]int, n)
	for i := range nxt {
		nxt[i] = make([]int, n)
	}

	nxt[0][n-1] = grid[0][0] + grid[0][n-1]

	for i := 1; i < m; i++ {
		for j1 := 0; j1 < n; j1++ {
			for j2 := 0; j2 < n; j2++ {
				best := -1
				for dj1 := j1 - 1; dj1 <= j1+1; dj1++ {
					for dj2 := j2 - 1; dj2 <= j2+1; dj2++ {
						if dj1 < 0 || dj1 >= n || dj2 < 0 || dj2 >= n || dp[dj1][dj2] == -1 {
							continue
						}
						tmp := dp[dj1][dj2]
						if j1 == j2 {
							tmp += grid[i][j1]
						} else {
							tmp += grid[i][j1] + grid[i][j2]
						}
						best = max(best, tmp)
					}
				}
				nxt[j1][j2] = best
			}
		}
		dp, nxt = nxt, dp
	}

	ans := 0
	for j1 := 0; j1 < n; j1++ {
		for j2 := 0; j2 < n; j2++ {
			ans = max(ans, dp[j1][j2])
		}
	}

	return ans
}
