package LC3349

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/14 下午10:33
 * @FilePath: has_increasing_subarrays
 * @Description: 顺序遍历记录前一个连续升序子段和当前连续升序子段
 */

func hasIncreasingSubarrays(nums []int, k int) bool {
	var j, n, pre, cur int
	n = len(nums)
	for i := 0; i < n; i++ {
		j = i
		for i+1 < n && nums[i] < nums[i+1] {
			i++
		}
		cur = i - j + 1
		if cur >= 2*k || (pre >= k && cur >= k) {
			return true
		}
		pre = cur
	}
	return false
}
