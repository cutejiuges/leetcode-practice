package lc3097

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	n := len(nums)
	bits := make([]int, 32)

	for left, right := 0, 0; right < n; right++ {
		for i := 0; i < 32; i++ {
			bits[i] += nums[right] >> i & 1
		}

		for left <= right && valueOf(bits) >= k {
			ans = min(ans, right-left+1)
			for i := 0; i < 32; i++ {
				bits[i] -= nums[left] >> i & 1
			}
			left++
		}
	}

	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}

func valueOf(bits []int) int {
	ans := 0
	n := len(bits)
	for i := 0; i < n; i++ {
		if bits[i] > 0 {
			ans |= 1 << i
		}
	}
	return ans
}
