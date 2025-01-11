package race420_q1

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/20 上午10:31
 * @FilePath: string_sequence
 * @Description:
 */

func stringSequence(target string) []string {
	ans := make([]string, 0)
	for _, ch := range target {
		if len(ans) == 0 {
			ans = append(ans, "a")
		} else {
			ans = append(ans, ans[len(ans)-1]+"a")
		}
		cur := ans[len(ans)-1]
		if int(ch) != int(cur[len(cur)-1]) {
			for i := 0; i < int(ch)-int('a'); i++ {
				ans = append(ans, cur[:len(cur)-1]+string(rune('a'+i+1)))
			}
		}
	}
	return ans
}
