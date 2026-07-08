package LC3756

var MOD = int(1e9 + 7)

func sumAndMultiply(s string, queries [][]int) []int {
	x, sum, cnt, base := initPrefixArray(s)
	length := len(queries)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		left, right := queries[i][0], queries[i][1]+1
		ss := sum[right] - sum[left]
		le := cnt[right] - cnt[left]
		xx := x[right] - x[left]*base[le]%int64(MOD) + int64(MOD)
		ans[i] = int(xx * ss % int64(MOD))
	}
	return ans
}

func initPrefixArray(s string) (x, sum, cnt, base []int64) {
	base = make([]int64, 100001)
	base[0] = int64(1)
	for i := 1; i < 100001; i++ {
		base[i] = base[i-1] * int64(10) % int64(MOD)
	}
	n := len(s)
	x, sum, cnt = make([]int64, n+1), make([]int64, n+1), make([]int64, n+1)
	for i := 0; i < n; i++ {
		cur := s[i] - '0'
		sum[i+1] = sum[i] + int64(cur)
		if cur > 0 {
			x[i+1] = (x[i]*int64(10) + int64(cur)) % int64(MOD)
			cnt[i+1] = cnt[i] + 1
		} else {
			x[i+1] = x[i]
			cnt[i+1] = cnt[i]
		}
	}
	return
}
