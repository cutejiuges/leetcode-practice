package LC2094

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/12 上午9:20
 * @FilePath: find_even_numbers
 * @Description:
 */

func findEvenNumbers(digits []int) []int {
	cnt := make([]int, 10)
	for _, digit := range digits {
		cnt[digit]++
	}
	var ans []int
	for i := 100; i < 999; i += 2 {
		if judge(i, cnt) {
			ans = append(ans, i)
		}
	}
	return ans
}

func judge(num int, cnt []int) bool {
	c := make([]int, 10)
	for x := num; x > 0; x /= 10 {
		cur := x % 10
		c[cur]++
		if c[cur] > cnt[cur] {
			return false
		}
	}
	return true
}
