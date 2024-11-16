package LC3240

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午9:04
 * @FilePath: min_flips
 * @Description:
 */

func minFlips(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0

	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			cnt1 := grid[i][j] + grid[m-i-1][j] + grid[i][n-j-1] + grid[m-i-1][n-j-1]
			ans += min(cnt1, 4-cnt1)
		}
	}

	diff, cnt1 := 0, 0
	if m&1 == 1 {
		for j := 0; j < n/2; j++ {
			if grid[m/2][j] != grid[m/2][n-j-1] {
				diff++
			} else {
				cnt1 += grid[m/2][j] * 2
			}
		}
	}
	if n&1 == 1 {
		for i := 0; i < m/2; i++ {
			if grid[i][n/2] != grid[m-i-1][n/2] {
				diff++
			} else {
				cnt1 += grid[i][n/2] * 2
			}
		}
	}
	if m&1 == 1 && n&1 == 1 {
		ans += grid[m/2][n/2]
	}

	if diff > 0 {
		ans += diff
	} else {
		ans += cnt1 % 4
	}
	return ans
}
