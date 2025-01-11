package LC3191

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/18 上午9:03
 * @FilePath: min_operations
 * @Description:
 */

func minOperations(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			if i > n-3 {
				return -1
			}

			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			ans++
		}
	}
	return ans
}
