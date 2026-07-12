package LC1331

import "sort"

func arrayRankTransform(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	sort.Ints(copyArr)

	idx := make(map[int]int)
	sameCnt := 0
	for i, num := range copyArr {
		if i > 0 && copyArr[i-1] == copyArr[i] {
			sameCnt++
		}
		if _, ok := idx[num]; !ok {
			idx[num] = i + 1 - sameCnt
		}
	}

	ans := make([]int, len(arr))
	for i, num := range arr {
		ans[i] = idx[num]
	}
	return ans
}
