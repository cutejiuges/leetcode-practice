package LC498

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/25 上午7:15
 * @FilePath: find_diagonal_order
 * @Description: lc498，对角线遍历，模拟，注意边界
 */

func findDiagonalOrder(mat [][]int) []int {
	row, col := len(mat), len(mat[0])
	cnt := row + col - 1
	var idx, m, n int
	res := make([]int, row*col)
	for i := 0; i < cnt; i++ {
		if i&1 == 0 { // 偶数个对角线向右上方
			for m >= 0 && n < col { // 没越界正常遍历
				res[idx] = mat[m][n]
				idx++
				m--
				n++
			}
			if n < col { // 处理行越界，直接移动到下一行
				m++
			} else { // 列越界，下移两行，前移一列
				m += 2
				n--
			}
		} else { // 奇数个对角线向左下方
			for m < row && n >= 0 { // 没越界正常遍历
				res[idx] = mat[m][n]
				idx++
				m++
				n--
			}
			if m < row {
				n++
			} else {
				n += 2
				m--
			}
		}
	}
	return res
}
