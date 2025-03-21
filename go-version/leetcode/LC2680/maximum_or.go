package LC2680

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/21 下午10:34
 * @FilePath: maximum_or
 * @Description: 位运算实现技巧
 */

func maximumOr(nums []int, k int) int64 {
	var allOr, fixed int // allOr表示所有数字按位或结果，fixed表示固定的1按位组成的结果
	for _, num := range nums {
		fixed |= allOr & num
		allOr |= num
	}

	var ans int64
	for _, num := range nums {
		ans = max(ans, int64((allOr^num)|fixed|num<<k))
	}
	return ans
}
