package lc47

import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	visited := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}

		for i, val := range nums {
			if visited[i] || i > 0 && !visited[i-1] && val == nums[i-1] {
				continue
			}
			perm = append(perm, nums[i])
			visited[i] = true
			backtrack(idx + 1)
			visited[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return ans
}
