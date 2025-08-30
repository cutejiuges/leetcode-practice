package LC36

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/30 上午11:31
 * @FilePath: is_valid_sudoku
 * @Description: 判断有效数独
 */

func isValidSudoku(board [][]byte) bool {
	var row, col [9][9]bool
	var box [3][3][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '1'
			if row[i][num] || col[j][num] || box[i/3][j/3][num] {
				return false
			}
			row[i][num] = true
			col[j][num] = true
			box[i/3][j/3][num] = true
		}
	}
	return true
}
