package LC3102

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/9 下午11:37
 * @FilePath: minimum_distance
 * @Description: 曼哈顿距离转换切比雪夫距离
 */

func minimumDistance(points [][]int) int {
	n := len(points)
	sx, sy := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		sx[i] = []int{x - y, i}
		sy[i] = []int{x + y, i}
	}
	sort.Slice(sx, func(i, j int) bool {
		return sx[i][0] < sx[j][0]
	})
	sort.Slice(sy, func(i, j int) bool {
		return sy[i][0] < sy[j][0]
	})

	maxVal1, maxVal2 := sx[n-1][0]-sx[0][0], sy[n-1][0]-sy[0][0]
	res := int(^uint(0) >> 1)
	if maxVal1 >= maxVal2 {
		i, j := sx[0][1], sx[n-1][1]
		res = min(res, max(remove(sx, i), remove(sy, i)))
		res = min(res, max(remove(sx, j), remove(sy, j)))
	} else {
		i, j := sy[0][1], sy[n-1][1]
		res = min(res, max(remove(sx, i), remove(sy, i)))
		res = min(res, max(remove(sx, j), remove(sy, j)))
	}
	return res
}

func remove(arr [][]int, i int) int {
	n := len(arr)
	if arr[0][1] == i {
		return arr[n-1][0] - arr[1][0]
	} else if arr[n-1][1] == i {
		return arr[n-2][0] - arr[0][0]
	} else {
		return arr[n-1][0] - arr[0][0]
	}
}
