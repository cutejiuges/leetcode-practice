package lc1561

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	n, round := len(piles), len(piles)/3
	ans, idx := 0, n-2
	for i := 0; i < round; i++ {
		ans += piles[idx]
		idx -= 2
	}
	return ans
}
