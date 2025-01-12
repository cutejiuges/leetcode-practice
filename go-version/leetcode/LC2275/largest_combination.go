package lc2275

import "math/bits"

func largestCombination(candidates []int) int {
	//取出数组中的最大值
	mx := candidates[0]
	for i := range candidates {
		mx = max(mx, candidates[i])
	}
	//最大值二进制有几位
	maxLen := bits.Len(uint(mx))

	//枚举二进制位，计算该位非零的条件下，最多可以取出多少个数
	//求出满足条件的最大值
	ans := 0
	for i := 0; i < maxLen; i++ {
		curCnt := 0
		for _, num := range candidates {
			curCnt += num >> i & 1
		}
		ans = max(ans, curCnt)
	}
	return ans
}
