package LC2145

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/21 上午8:59
 * @FilePath: number_of_arrays
 * @Description: 前缀和数组最值
 */

func numberOfArrays(differences []int, lower int, upper int) int {
	var base, mx, mn int
	for _, diff := range differences {
		base += diff
		mx = max(mx, base)
		mn = min(mn, base)
		if mx-mn > upper-lower {
			return 0
		}
	}
	return 1 + (upper - lower) - (mx - mn)
}
