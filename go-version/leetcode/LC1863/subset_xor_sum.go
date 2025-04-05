package LC1863

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/5 下午3:36
 * @FilePath: subset_xor_sum
 * @Description: 通过深度优先搜索遍历全部子集
 */

func subsetXORSum(nums []int) int {
	ans := 0
	n := len(nums)
	var dfs func(int, int, []int)

	dfs = func(val int, idx int, arr []int) {
		if idx == n { //递归终止条件
			ans += val
			return
		}
		// 选择当前数字
		dfs(val^arr[idx], idx+1, arr)
		// 跳过当前数字
		dfs(val, idx+1, arr)
	}

	dfs(0, 0, nums)
	return ans
}
