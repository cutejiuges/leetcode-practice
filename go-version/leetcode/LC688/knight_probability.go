package LC688

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午7:45
 * @FilePath: knight_probability
 * @Description: 动态规划问题
 */

func knightProbability(n int, k int, row int, column int) float64 {
	var dirs = [][]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}
	//定义dp[step][i][j]是棋子从(i,j)位置出发，走了step步之后仍然留在棋盘上的概率，
	//当(i,j)不在棋盘上时，dp[step][i][j]=0，当(i,j)在棋盘上且step=0时，dp[step][i][j]=1
	//对于一般情况，dp[step][i][j] = (1/8)*sum(dp[step-1][i+dx][j+dy])

	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if i+d[0] >= 0 && i+d[0] < n && j+d[1] >= 0 && j+d[1] < n {
							dp[step][i][j] += 0.125 * dp[step-1][i+d[0]][j+d[1]]
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
