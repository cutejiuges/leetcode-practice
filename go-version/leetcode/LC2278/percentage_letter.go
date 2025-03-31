package LC2278

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/31 下午11:25
 * @FilePath: percentage_letter
 * @Description: 统计字符数量即可
 */

func percentageLetter(s string, letter byte) int {
	cnt, n := 0, len(s)
	for _, ch := range s {
		if byte(ch) == letter {
			cnt++
		}
	}
	return cnt * 100 / n
}
