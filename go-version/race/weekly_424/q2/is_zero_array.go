package main

import "fmt"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/17 上午10:50
 * @FilePath: is_zero_array
 * @Description:
 */

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	//差分数组
	diffArray := make([]int, n)
	for _, query := range queries {
		diffArray[query[0]]++
		if query[1] < n-1 {
			diffArray[query[1]+1]--
		}
	}

	cntArray := make([]int, n)
	cntArray[0] = diffArray[0]
	for i := 1; i < n; i++ {
		cntArray[i] = cntArray[i-1] + diffArray[i]
	}

	for i, num := range nums {
		if num > cntArray[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}}))
	fmt.Println(isZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}))
	fmt.Println(isZeroArray([]int{4, 6}, [][]int{{0, 0}, {0, 1}, {1, 1}, {0, 0}, {1, 1}, {1, 1}, {0, 0}, {1, 1}, {1, 1}, {1, 1}, {0, 0}, {1, 1}, {0, 1}, {0, 0}, {1, 1}}))
}
