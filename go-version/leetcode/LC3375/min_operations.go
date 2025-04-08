package LC3375

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/9 上午12:16
 * @FilePath: min_operations
 * @Description: 本质是寻找不同的数字个数
 */

func minOperations(nums []int, k int) int {
	st := map[int]struct{}{}
	for _, m := range nums {
		if m < k {
			return -1
		} else if m > k {
			st[m] = struct{}{}
		}
	}
	return len(st)
}
