package lc3095

import "math"

// 使用滑动窗口来实现，但是因为或运算没有逆运算，因此只能统计每一位出现1的次数
// 如果某一位出现1的次数不是0那么最终或起来这一位就是1
// 将右端点的位统计到窗口中，左端点的位从窗口中拿掉，计算窗口中留存的位能否满足k要求
func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	bits := make([]int, 32)
	res := math.MaxInt

	for left, right := 0, 0; right < n; right++ {
		for i := 0; i < 32; i++ {
			bits[i] += nums[right] >> i & 1
		}

		for left <= right && value(bits) >= k {
			res = min(res, right-left+1)
			for i := 0; i < 32; i++ {
				bits[i] -= nums[left] >> i & 1
			}
			left++
		}
	}

	if res == math.MaxInt {
		res = -1
	}
	return res
}

// 通过位计算实际的值
func value(bits []int) int {
	ans, n := 0, len(bits)
	for i := 0; i < n; i++ {
		if bits[i] > 0 {
			ans |= 1 << i
		}
	}
	return ans
}
