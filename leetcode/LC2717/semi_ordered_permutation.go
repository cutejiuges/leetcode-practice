package LC2717

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/11 下午8:52
 * @FilePath: semi_ordered_permutation
 * @Description:
 */

func semiOrderedPermutation(nums []int) int {
	first, last := 0, 0
	n := len(nums)
	for i := range nums {
		if nums[i] == 1 {
			first = i
		}
		if nums[i] == n {
			last = i
		}
	}
	ans := first + n - last - 1
	if first > last {
		ans -= 1
	}
	return ans
}
