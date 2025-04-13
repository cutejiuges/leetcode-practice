package LC1922

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/13 下午6:48
 * @FilePath: count_good_numbers
 * @Description: 快速幂
 */

func countGoodNumbers(n int64) int {
	mod := int64(1e9 + 7)
	return int(quickPower(5, (n+1)/2, mod) * quickPower(4, n/2, mod) % mod)
}

func quickPower(a, b, mod int64) int64 {
	ans := int64(1)
	for b > 0 {
		if b&1 == 1 {
			ans = ans * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return ans
}
