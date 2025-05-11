package q2

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/11 上午10:50
 * @FilePath: can_partition_grid
 * @Description:
 */

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	var totalSum int
	rowSum, colSum := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			totalSum += grid[i][j]
			rowSum[i] += grid[i][j]
			colSum[j] += grid[i][j]
		}
	}
	if totalSum&1 == 1 {
		return false
	}

	target := totalSum >> 1
	// 尝试水平分割
	if hRes := judge(rowSum, target); hRes {
		return true
	}
	// 尝试竖直分割
	if vRes := judge(colSum, target); vRes {
		return true
	}
	return false
}

func judge(sum []int, target int) bool {
	cur, n := 0, 0
	for i := 0; i < n; i++ {
		cur += sum[i]
		if cur == target {
			return true
		} else if cur > target {
			return false
		}
	}
	return false
}
