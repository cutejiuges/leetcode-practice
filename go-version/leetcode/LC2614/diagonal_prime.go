package LC2614

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/18 下午11:02
 * @FilePath: diagonal_prime
 * @Description: 艾氏筛进行素数判别
 */

func diagonalPrime(nums [][]int) int {
	primes := initPrimes()
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if primes[nums[i][i]] {
			ans = max(ans, nums[i][i])
		}
		if primes[nums[i][n-i-1]] {
			ans = max(ans, nums[i][n-i-1])
		}
	}
	return ans
}

func initPrimes() []bool {
	n := int(4*1e6 + 1)
	primes := make([]bool, n)
	for i := range primes {
		primes[i] = true
	}
	primes[0], primes[1] = false, false
	for i := 2; i*i < n; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	return primes
}
