package LC2799

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/24 下午11:12
 * @FilePath: count_complete_subarrays
 * @Description: 滑动窗口
 */

func countCompleteSubarrays(nums []int) int {
	distinct := differentNums(nums)
	ans, left := 0, 0
	mp := make(map[int]int)
	for i := range nums {
		mp[nums[i]]++
		for len(mp) >= distinct {
			mp[nums[left]]--
			if mp[nums[left]] == 0 {
				delete(mp, nums[left])
			}
			left++
		}
		ans += left
	}
	return ans
}

func differentNums(nums []int) int {
	mp := make(map[int]struct{})
	for _, num := range nums {
		mp[num] = struct{}{}
	}
	return len(mp)
}
