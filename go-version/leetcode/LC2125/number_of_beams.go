package LC2125

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/27 下午11:26
 * @FilePath: number_of_beams
 * @Description: 按行遍历即可
 */

func numberOfBeams(bank []string) int {
	var ans, pre int
	for _, line := range bank {
		var cnt int
		for _, ch := range line {
			cnt += int(ch - '0')
		}
		if cnt > 0 {
			ans += cnt * pre
			pre = cnt
		}
	}
	return ans
}
