package LC551

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/18 下午7:58
 * @FilePath: check_record
 * @Description: 字符串遍历模拟
 */

func checkRecord(s string) bool {
	cntA, cntL := 0, 0
	for _, ch := range s {
		if ch == 'A' {
			cntA++
			if cntA >= 2 {
				return false
			}
		}

		if ch == 'L' {
			cntL++
			if cntL >= 3 {
				return false
			}
		} else {
			cntL = 0
		}
	}
	return true
}
