package LC594

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/30 下午10:54
 * @FilePath: find_LHS
 * @Description:
 */

func findLHS(nums []int) int {
	var ans int
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	for val, cnt := range mp {
		if cntNext, ok := mp[val+1]; ok {
			ans = max(ans, cnt+cntNext)
		}
	}
	return ans
}
