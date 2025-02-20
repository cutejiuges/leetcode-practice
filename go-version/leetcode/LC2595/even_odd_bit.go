package lc2595

func evenOddBit(n int) []int {
	ans := make([]int, 2)
	pos := 0
	for n > 0 {
		ans[pos] += n & 1
		n >>= 1
		pos ^= 1
	}
	return ans
}
