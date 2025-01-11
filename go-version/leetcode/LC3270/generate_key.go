package lc3270

func generateKey(num1 int, num2 int, num3 int) int {
	pow, ans := 1, 0
	for num1 != 0 && num2 != 0 && num3 != 0 {
		remainder1, remainder2, remainder3 := num1%10, num2%10, num3%10
		ans += pow * (min(remainder1, remainder2, remainder3))
		num1, num2, num3 = num1/10, num2/10, num3/10
		pow *= 10
	}
	return ans
}
