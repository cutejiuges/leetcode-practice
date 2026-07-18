package LC1979

import "math"

func findGCD(nums []int) int {
	mn, mx := math.MaxInt, math.MinInt
	for _, num := range nums {
		if num <= mn {
			mn = num
		}
		if num >= mx {
			mx = num
		}
	}
	return gcd(mn, mx)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
