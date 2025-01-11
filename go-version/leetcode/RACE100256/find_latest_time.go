package RACE100256

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午10:32
 * @FilePath: find_latest_time
 * @Description: 替换字符串的最晚时间
 */
func findLatestTime(s string) string {
	ss := []rune(s)
	for i := range ss {
		if ss[i] == '?' {
			if i == 0 {
				if ss[1] <= '1' || ss[1] == '?' {
					ss[i] = '1'
				} else {
					ss[i] = '0'
				}
			} else if i == 1 {
				if ss[0] == '0' {
					ss[i] = '9'
				} else {
					ss[i] = '1'
				}
			} else if i == 3 {
				ss[i] = '5'
			} else {
				ss[i] = '9'
			}
		}
	}
	return string(ss)
}
