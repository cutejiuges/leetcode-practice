/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-24 12:11:08
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-24 13:01:33
 * @FilePath: /leetcode-practice/leetcode/LC2684/max_moves.go
 * @Description: 最大移动次数问题，考虑使用bfs或者dfs来实现
 */
package lc2684

import "cutejiuge/leetcode-practice/math_util"

func maxMoves(grid [][]int) int {
	bfs(grid)
	return dfs(grid)
}

func bfs(grid [][]int) int {
	height, width := len(grid), len(grid[0])
	queue := make(map[int]bool)
	for i := range grid {
		queue[i] = true
	}

	for col := 1; col < width; col++ {
		p := make(map[int]bool)
		for i := range queue {
			for row := i - 1; row <= i+1; row++ {
				if row >= 0 && row < height && grid[row][col] > grid[i][col-1] {
					p[row] = true
				}
			}
		}
		queue = p
		if len(queue) == 0 {
			return col - 1
		}
	}
	return width - 1
}

func dfs(grid [][]int) int {
	height, width := len(grid), len(grid[0])
	dp := make([][]int, height)
	for i := range dp {
		dp[i] = make([]int, width)
	}

	for col := width - 2; col >= 0; col-- {
		for row := 0; row < height; row++ {
			mx := 0
			if grid[row][col] < grid[row][col+1] {
				mx = math_util.Max(mx, dp[row][col+1]+1)
			}
			if row+1 < height && grid[row][col] < grid[row+1][col+1] {
				mx = math_util.Max(mx, dp[row+1][col+1]+1)
			}
			if row-1 >= 0 && grid[row][col] < grid[row-1][col+1] {
				mx = math_util.Max(mx, dp[row-1][col+1]+1)
			}
			dp[row][col] = mx
		}
	}

	res := dp[0][0]
	for i := range dp {
		res = math_util.Max(res, dp[i][0])
	}
	return res
}
