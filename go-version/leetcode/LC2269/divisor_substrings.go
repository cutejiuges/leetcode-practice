package LC2269

import "strconv"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/10 下午11:26
 * @FilePath: divisor_substrings
 * @Description: 枚举每一个字串即可
 */

func divisorSubstrings(num, k int) int {
	ss := strconv.Itoa(num)
	n := len(ss)
	ans := 0
	for i := 0; i <= n-k; i++ {
		t, _ := strconv.Atoi(ss[i : i+k])
		if t != 0 && num%t == 0 {
			ans++
		}
	}
	return ans
}
