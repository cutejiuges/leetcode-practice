package LC3001

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/5 下午10:52
 * @FilePath: min_moves_to_capture_the_queen
 * @Description: 分类讨论模拟即可
 */

func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	//如果车和后在同一行
	if a == e {
		// 且中间没有象，那么只需要一次
		if a != c || d < min(b, f) || d > max(b, f) {
			return 1
		}
	}
	//如果车和后在同一列
	if b == f {
		//且中间没有象，也只需要一次
		if b != d || c < min(a, e) || c > max(a, e) {
			return 1
		}
	}
	//如果象和后在同一个对角线
	if abs(c-e) == abs(d-f) {
		if (c-e)*(b-f) != (a-e)*(d-f) || a < min(c, e) || a > max(c, e) {
			return 1
		}
	}
	return 2
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
