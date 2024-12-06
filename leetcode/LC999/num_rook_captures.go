package LC999

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/6 下午10:36
 * @FilePath: num_rook_captures
 * @Description: 直接模拟
 */

func numRookCaptures(board [][]byte) int {
	//寻找白车的位置
	x, y := 0, 0
	for i := 0; i < 64; i++ {
		if board[i/8][i%8] == 'R' {
			x, y = i/8, i%8
		}
	}

	//从4个方向出发，直接模拟可以找到的位置
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	ans := 0
	//遍历可移动的四个方向
	for i := 0; i < 4; i++ {
		//模拟一步步前进
		for step := 0; ; step++ {
			curx, cury := x+dx[i]*step, y+dy[i]*step
			if curx < 0 || curx >= 8 || cury < 0 || cury >= 8 || board[curx][cury] == 'B' {
				break
			}
			if board[curx][cury] == 'p' {
				ans++
				break
			}
		}
	}
	return ans
}
