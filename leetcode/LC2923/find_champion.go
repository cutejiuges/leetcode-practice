package LC2923

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/12 上午8:42
 * @FilePath: find_champion
 * @Description: 简单题，直接模拟即可
 */

func findChampion(grid [][]int) int {
	n := len(grid)
	for i := range grid {
		if sum(grid[i]) >= n-1 {
			return i
		}
	}
	return -1
}

func sum(nums []int) int {
	ans := 0
	for i := range nums {
		ans += nums[i]
	}
	return ans
}

func findChampion2(grid [][]int) int {
	n := len(grid)
	champion := 0
	for i := 1; i < n; i++ {
		if grid[i][champion] == 1 {
			champion = i
		}
	}
	return champion
}
