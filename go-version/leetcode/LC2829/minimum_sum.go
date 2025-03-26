package LC2829

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/26 上午8:33
 * @FilePath: minimum_sum
 * @Description: 使用贪心+等差数列
 */

func minimumSum(n int, k int) int {
	leftSum := sumArithmeticSequence(1, 1, min(n, k>>1))
	rightSum := sumArithmeticSequence(k, 1, n-(k>>1))
	if n <= (k >> 1) {
		return leftSum
	}
	return leftSum + rightSum
}

func sumArithmeticSequence(a1, d, n int) int {
	return n*a1 + ((n - 1) * n * d >> 1)
}
