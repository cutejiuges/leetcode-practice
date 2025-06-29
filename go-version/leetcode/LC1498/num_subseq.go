package LC1498

import "slices"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/29 下午10:51
 * @FilePath: num_subseq
 * @Description:
 */

const MOD = 1000000007

var pow2 = [100000]int{1}

func init() {
	for i := 1; i < len(pow2); i++ {
		pow2[i] = pow2[i-1] * 2 % MOD
	}
}

func numSubseq(nums []int, target int) int {
	ans := 0
	slices.Sort(nums)
	low, high := 0, len(nums)-1
	for low <= high {
		if nums[low]+nums[high] <= target {
			ans += pow2[high-low]
			low++
		} else {
			high--
		}
	}
	return ans % MOD
}
