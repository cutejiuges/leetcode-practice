package LC3532

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	collect := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > maxDiff {
			collect[i] = i
		} else {
			collect[i] = collect[i-1]
		}
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = collect[q[0]] == collect[q[1]]
	}
	return ans
}
