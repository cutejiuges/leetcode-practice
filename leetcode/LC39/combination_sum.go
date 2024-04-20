package LC39

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/20 下午6:17
 * @FilePath: combination_sum
 * @Description: 组合总数，经典回溯问题
 */

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(int, int)
	dfs = func(target int, idx int) {
		//如果遍历完了整个数组，直接退出
		if idx == len(candidates) {
			return
		}
		//如果当前的目标值已经是0,在答案中追加当前组合并退出
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		//直接跳过当前数字
		dfs(target, idx+1)
		//选择当前数字
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			//恢复现场
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}
