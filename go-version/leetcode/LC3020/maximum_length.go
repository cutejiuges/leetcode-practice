package LC3020

func maximumLength(nums []int) int {
	cnt := make(map[int64]int)
	for _, num := range nums {
		cnt[int64(num)]++
	}

	// ans至少是1的数量，向下取奇数
	oneCnt := cnt[1]
	ans := oneCnt
	if oneCnt&1 == 0 {
		ans--
	}

	delete(cnt, 1)

	for num := range cnt {
		res := 0
		x := num
		for cnt[x] > 1 {
			res += 2
			x *= x
		}
		if _, ok := cnt[x]; ok {
			ans = max(ans, res+1)
		} else {
			ans = max(ans, res-1)
		}
	}
	return ans
}
