package LC1590

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/11/30 下午8:45
 * @FilePath: min_subarray
 * @Description:
 */

func minSubarray(nums []int, p int) int {
	var sum int64
	for _, num := range nums {
		sum += int64(num)
	}
	remain := sum % int64(p)
	if remain == 0 {
		return 0
	}

	n := len(nums)
	ans := n
	s := 0
	last := map[int]int{s: -1}
	for i := 0; i < n; i++ {
		s = (s + nums[i]) % p
		last[s] = i
		var j int
		if idx, ok := last[(s-int(remain)+p)%p]; ok {
			j = idx
		} else {
			j = -n
		}
		ans = min(ans, i-j)
	}
	if ans == n {
		return -1
	}
	return ans
}
