package LC2597

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/7 上午9:18
 * @FilePath: beautiful_subsets
 * @Description: 美丽子集，暴力回溯
 */

func beautifulSubsets(nums []int, k int) int {
	mp := map[int]int{}
	ans := -1
	var backtrack func(int, []int, int, map[int]int)
	backtrack = func(idx int, nums []int, k int, m map[int]int) {
		//枚举终结
		if idx == len(nums) {
			ans++
			return
		}
		//不选择当前直接跳过
		backtrack(idx+1, nums, k, m)
		//判断能选就选中
		cur := nums[idx]
		if m[cur-k] == 0 && m[cur+k] == 0 { //可选
			m[cur]++
			backtrack(idx+1, nums, k, m)
			//恢复现场
			m[cur]--
		}
	}
	backtrack(0, nums, k, mp)
	return ans
}
