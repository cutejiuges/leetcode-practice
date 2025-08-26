package LC3000

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/26 上午8:19
 * @FilePath: area_of_max_diagonal
 * @Description: 一次遍历
 */

func areaOfMaxDiagonal(dimensions [][]int) int {
	var maxArea, maxLength int
	for _, dimension := range dimensions {
		area := dimension[0] * dimension[1]
		length := dimension[0]*dimension[0] + dimension[1]*dimension[1]
		if length > maxLength {
			maxLength = length
			maxArea = area
		} else if length == maxLength {
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}
