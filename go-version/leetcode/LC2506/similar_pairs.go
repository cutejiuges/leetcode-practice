package LC2506

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/22 下午9:40
 * @FilePath: similar_pairs
 * @Description: 实现相似字符串数对
 */

func similarPairs(words []string) int {
	ans := 0
	mp := make(map[int]int)
	for _, ss := range words {
		mask := 0
		for _, c := range ss {
			mask |= 1 << (c - 'a')
		}
		ans += mp[mask]
		mp[mask]++
	}
	return ans
}
