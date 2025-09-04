package LC3516

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/4 下午11:38
 * @FilePath: find_closest
 * @Description: 模拟题直接签到
 */

func findClosest(x, y, z int) int {
	a, b := abs(x-z), abs(y-z)
	if a == b {
		return 0
	} else if a < b {
		return 1
	} else {
		return 2
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
