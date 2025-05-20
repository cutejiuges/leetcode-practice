package LC3355

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/20 下午11:30
 * @FilePath: is_zero_array
 * @Description:
 */

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diffArray := make([]int, n)
	for _, query := range queries {
		diffArray[query[0]]++
		if query[1] < n-1 {
			diffArray[query[1]+1]--
		}
	}

	sumDiff := 0
	for i := 0; i < n; i++ {
		sumDiff += diffArray[i]
		if nums[i] > sumDiff {
			return false
		}
	}
	return true
}
