package LC2537

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/16 上午8:30
 * @FilePath: count_good
 * @Description: 滑动窗口
 */

func countGood(nums []int, k int) int64 {
	cnt := make(map[int]int)
	var ans int64
	left, pairs := 0, 0
	for _, num := range nums {
		pairs += cnt[num]
		cnt[num]++
		for pairs >= k {
			pairs -= cnt[nums[left]] - 1
			cnt[nums[left]]--
			left++
		}
		ans += int64(left)
	}
	return ans
}
