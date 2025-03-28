package LC2716

import "math/bits"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/28 下午11:23
 * @FilePath: minimized_string_length
 * @Description: 统计字符串中不同字符数量即可
 */

func minimizedStringLength(s string) int {
	mask := uint(0)
	for _, ch := range s {
		mask |= 1 << (ch - 'a')
	}
	return bits.OnesCount(mask)
}
