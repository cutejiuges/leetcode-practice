package LC3396

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/8 下午11:46
 * @FilePath: minimum_operations
 * @Description: 倒序遍历进行检查
 */

func minimumOperations(nums []int) int {
	set := map[int]struct{}{}
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if _, ok := set[nums[i]]; ok {
			return i/3 + 1
		}
		set[nums[i]] = struct{}{}
	}
	return 0
}
