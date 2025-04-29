package LC1295

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/30 上午2:11
 * @FilePath: find_numbers
 * @Description: 辗转相除计算位数再累加即可
 */

func findNumbers(nums []int) int {
	var ans int
	for _, num := range nums {
		if count10Bits(num)&1 == 0 {
			ans++
		}
	}
	return ans
}

func count10Bits(num int) int {
	x, ans := num, 0
	for x > 0 {
		x /= 10
		ans++
	}
	return ans
}
