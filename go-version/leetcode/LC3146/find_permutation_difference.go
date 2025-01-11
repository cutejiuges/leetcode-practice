package LC3146

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/24 上午1:06
 * @FilePath: find_permutation_difference
 * @Description: 考虑hash存储位置信息即可
 */

func findPermutationDifference(s, t string) int {
	pos := [26]int{}
	for idx, ch := range s {
		pos[ch-'a'] = idx
	}

	ans := 0
	for idx, ch := range t {
		ans += abs(idx - pos[ch-'a'])
	}
	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
