package LC2598

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/16 下午11:38
 * @FilePath: find_smallest_integer
 * @Description: 同余计数
 */

func findSmallestInteger(nums []int, value int) int {
	mp := make([]int, value)
	for _, num := range nums {
		v := (num%value + value) % value
		mp[v]++
	}

	var mex int
	for mp[mex%value] > 0 {
		mp[mex%value]--
		mex++
	}
	return mex
}
