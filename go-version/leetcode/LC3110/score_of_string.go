package LC3110

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/15 下午9:29
 * @FilePath: score_of_string
 * @Description: 直接遍历模拟即可
 */

func scoreOfString(s string) int {
	var ans int
	n := len(s)
	for i := 0; i < n-1; i++ {
		ans += abs(int(s[i]) - int(s[i+1]))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
