package LC2918

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/10 下午8:39
 * @FilePath: min_sum
 * @Description: 最小和，遍历 + 分类
 */

func minSum(nums1, nums2 []int) int64 {
	sum1, cntZero1 := calculate(nums1)
	sum2, cntZero2 := calculate(nums2)
	if sum1 > sum2 && cntZero2 == 0 {
		return -1
	}
	if sum2 > sum1 && cntZero1 == 0 {
		return -1
	}
	return max(sum1, sum2)
}

func calculate(nums []int) (int64, int64) {
	var sum, cnt0 int64
	for _, num := range nums {
		sum += int64(num)
		if num == 0 {
			sum++
			cnt0++
		}
	}
	return sum, cnt0
}
