package LC3024

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/19 上午9:17
 * @FilePath: triangle_type
 * @Description: 排序分类判断即可
 */

func triangleType(nums []int) string {
	sort.Ints(nums)
	a, b, c := nums[0], nums[1], nums[2]
	if a+b <= c {
		return "none"
	}
	if a == b && b == c {
		return "equilateral"
	}
	if a == b || b == c || c == a {
		return "isosceles"
	}
	return "scalene"
}
