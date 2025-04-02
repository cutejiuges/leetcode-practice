package LC2873

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/2 下午11:01
 * @FilePath: maximum_triplet_value
 * @Description:
 */

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	sufMax := make([]int, n+1)
	for i := n - 1; i > 1; i-- {
		sufMax[i] = max(nums[i], sufMax[i+1])
	}
	var ans int64
	preMax := nums[0]
	for i := 1; i < n-1; i++ {
		ans = max(ans, int64((preMax-nums[i])*sufMax[i+1]))
		preMax = max(preMax, nums[i])
	}
	return ans
}
