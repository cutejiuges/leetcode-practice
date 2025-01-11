package main

import "fmt"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/24 上午10:31
 * @FilePath: minimum_sum_subarray
 * @Description:
 */

func minimumSumSubarray(nums []int, l int, r int) int {
	minSum := 0x3f3f3f3f

	for length := l; length <= r; length++ {
		for i := 0; i <= len(nums)-length; i++ {
			subSum := 0
			for j := i; j < i+length; j++ {
				subSum += nums[j]
			}
			if subSum > 0 && subSum < minSum {
				minSum = subSum
			}
		}
	}
	if minSum == 0x3f3f3f3f {
		return -1
	}
	return minSum
}

func main() {
	fmt.Println(minimumSumSubarray([]int{3, -2, 1, 4}, 2, 3))
	fmt.Println(minimumSumSubarray([]int{-2, 2, -3, 1}, 2, 3))
	fmt.Println(minimumSumSubarray([]int{1, 2, 3, 4}, 2, 4))
}
