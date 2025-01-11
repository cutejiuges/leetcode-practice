package LC419

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/11 上午12:07
 * @FilePath: count_battleships
 * @Description:
 */

func countBattleships(board [][]byte) int {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}
	row, col := len(board), len(board[0])
	ans := 0

	//遍历战舰群
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'X' { //如果是一个战舰
				board[i][j] = '.'                                    //防止重复计算
				for k := i + 1; k < row && board[k][j] == 'X'; k++ { //找同一行上相邻x，也是战舰的一部分
					board[k][j] = '.' //防止重复计算
				}
				for k := j + 1; k < col && board[i][k] == 'X'; k++ { //找同一列上相邻x，也是战舰的一部分
					board[i][k] = '.' //防止重复计算
				}
				ans++
			}
		}
	}
	return ans
}
