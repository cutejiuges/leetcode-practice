package LC904

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/4 下午11:22
 * @FilePath: total_fruit
 * @Description: 滑动窗口
 */

func totalFruit(fruits []int) int {
	var left, ans int
	mp := make(map[int]int)
	for i := 0; i < len(fruits); i++ {
		mp[fruits[i]]++
		for len(mp) > 2 {
			mp[fruits[left]]--
			if mp[fruits[left]] == 0 {
				delete(mp, fruits[left])
			}
			left++
		}
		ans = max(ans, i-left+1)
	}
	return ans
}
