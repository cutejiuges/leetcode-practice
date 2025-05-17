package LC75

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/17 下午9:57
 * @FilePath: sort_colors
 * @Description: 插入覆盖修改
 */

func sortColors(nums []int) {
	p0, p1 := 0, 0
	n := len(nums)
	for i := 0; i < n; i++ {
		x := nums[i]
		nums[i] = 2
		if x <= 1 {
			nums[p1] = 1
			p1++
		}
		if x == 0 {
			nums[p0] = 0
			p0++
		}
	}
}
