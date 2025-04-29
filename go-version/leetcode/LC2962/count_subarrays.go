package LC2962

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/29 下午11:13
 * @FilePath: count_subarrays
 * @Description: 滑动窗口，越长越合法
 */

func countSubarrays(nums []int, k int) int64 {
	var ans int64
	var cntMax, left, mx int
	for _, num := range nums {
		mx = max(mx, num)
	}

	for i := range nums {
		if nums[i] == mx {
			cntMax++
		}
		for cntMax >= k {
			if nums[left] == mx {
				cntMax--
			}
			left++
		}
		ans += int64(left)
	}
	return ans
}
