package LC3307

import "math/bits"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/4 下午11:27
 * @FilePath: kth_character
 * @Description: 向前回溯t次
 */

func kthCharacter(k int64, operations []int) byte {
	var ans, t int
	for k != 1 {
		t = bits.Len64(uint64(k)) - 1
		if (int64(1) << t) == k {
			t--
		}
		k = k - (int64(1) << t)
		if operations[t] != 0 {
			ans++
		}
	}
	return byte('a' + (ans % 26))
}
