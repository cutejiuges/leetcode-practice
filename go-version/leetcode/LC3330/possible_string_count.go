package LC3330

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/1 上午8:47
 * @FilePath: possible_string_count
 * @Description: 遍历寻找重复字符数
 */

func possibleStringCount(word string) int {
	ans := 1
	length := len(word)
	for i := 1; i < length; i++ {
		if word[i] == word[i-1] {
			ans++
		}
	}
	return ans
}
