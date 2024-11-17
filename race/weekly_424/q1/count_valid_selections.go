package main

import "fmt"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/17 上午10:30
 * @FilePath: count_valid_selections
 * @Description:
 */

func countValidSelections(nums []int) int {
	n := len(nums)
	sumArray := make([]int, n)
	sumArray[0] = nums[0]
	for i := 1; i < n; i++ {
		sumArray[i] = sumArray[i-1] + nums[i]
	}

	ans := 0
	for i, num := range nums {
		if num == 0 {
			if sumArray[i] == sumArray[n-1]-sumArray[i] {
				ans += 2
			} else if sumArray[i]*2-sumArray[n-1] == 1 || sumArray[i]*2-sumArray[n-1] == -1 {
				ans += 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 3}))
	fmt.Println(countValidSelections([]int{2, 3, 4, 0, 4, 1, 0}))
	fmt.Println(countValidSelections([]int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7}))
}
