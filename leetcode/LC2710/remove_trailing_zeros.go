package LC2710

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/29 下午8:38
 * @FilePath: remove_trailing_zeros
 * @Description: 字符串操作，直接遍历即可
 */

func removeTrailingZeros(num string) string {
	n := len(num) - 1
	for ; n >= 0 && num[n] == '0'; n-- {
	}
	return num[:n+1]
}
