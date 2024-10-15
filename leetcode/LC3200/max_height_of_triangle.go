package LC3200

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/15 下午11:14
 * @FilePath: max_height_of_triangle
 * @Description:
 */

func maxHeightOfTriangle(red int, blue int) int {
	return max(calc(red, blue), calc(blue, red))
}

func calc(x, y int) int {
	costX, costY := 1, 2
	for i := 0; ; i++ {
		if i%2 == 0 {
			if x < costX {
				return i
			}
			x -= costX
			costX += 2
		} else {
			if y < costY {
				return i
			}
			y -= costY
			costY += 2
		}
	}
}
