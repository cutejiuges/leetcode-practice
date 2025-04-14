package LC1534

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/14 上午9:05
 * @FilePath: count_good_triplets
 * @Description: 暴力遍历
 */

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var ans int
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					ans++
				}
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
