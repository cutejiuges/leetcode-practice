package LC1338

import (
	"sort"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/15 下午8:16
 * @FilePath: min_set_size
 * @Description: 贪心解决
 */

func minSetSize(arr []int) int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	var cnt []int
	for _, c := range freq {
		cnt = append(cnt, c)
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})

	deletedCnt := 0
	n := len(arr)
	for i := range cnt {
		deletedCnt += cnt[i]
		if deletedCnt >= n/2 {
			return i + 1
		}
	}
	return -1
}
