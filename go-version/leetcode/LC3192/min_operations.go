package LC3192

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/19 上午10:26
 * @FilePath: min_operations
 * @Description:
 */

func minOperations(nums []int) int {
	operations := 0
	for _, num := range nums {
		if num == (operations & 1) {
			operations++
		}
	}
	return operations
}
