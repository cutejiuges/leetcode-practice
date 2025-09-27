package LC812

import "math"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/27 下午10:46
 * @FilePath: largest_triangle_area
 * @Description:
 */

func largestTriangleArea(points [][]int) float64 {
	var maxArea float64
	var n = len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := 0.5 * math.Abs(
					float64(points[i][0]*(points[j][1]-points[k][1])+
						points[j][0]*(points[k][1]-points[i][1])+
						points[k][0]*(points[i][1]-points[j][1])),
				)
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}
