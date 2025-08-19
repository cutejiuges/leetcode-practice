package LC2348

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/19 下午11:26
 * @FilePath: zero_filled_subarray
 * @Description:
 */

func zeroFilledSubarray(nums []int) int64 {
	var cnt, ans int64
	for _, num := range nums {
		if num == 0 {
			cnt++
			ans += cnt
		} else {
			cnt = 0
		}
	}
	return ans
}
