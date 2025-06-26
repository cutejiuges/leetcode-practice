package LC2311

import (
	"math/bits"
	"strconv"
	"strings"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/26 下午11:05
 * @FilePath: longest_subsequence
 * @Description:
 */

func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if n < m {
		// 可以全选
		return n
	}
	ans := m //后缀长度
	sufVal, _ := strconv.ParseInt(s[n-m:], 2, 0)
	if int(sufVal) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0")
}
