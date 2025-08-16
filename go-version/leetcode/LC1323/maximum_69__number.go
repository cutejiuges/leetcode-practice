package LC1323

import "strconv"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/16 下午10:04
 * @FilePath: maximum_69__number
 * @Description:
 */

func maximum69Number(num int) int {
	ss := []byte(strconv.Itoa(num))
	for i := 0; i < len(ss); i++ {
		if ss[i] == '6' {
			ss[i] = '9'
			break
		}
	}
	num, _ = strconv.Atoi(string(ss))
	return num
}
