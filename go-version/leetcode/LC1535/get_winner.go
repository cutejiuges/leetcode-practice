package LC1535

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/29 下午8:56
 * @FilePath: get_winner
 * @Description: 模拟擂台赛机制，新来一个单挑一个
 */

func getWinner(arr []int, k int) int {
	n := len(arr)
	if n == 0 || k == 0 {
		return 0
	}

	//先假定数组的首个元素就是获胜者，默认是第0回合
	win, turn := arr[0], 0
	for i := 1; i < n && k > turn; i++ {
		//如果擂台赛被击败，获胜者换人，获胜回合清空
		if arr[i] > win {
			win = arr[i]
			turn = 0
		}
		//获胜回合加一个，换人之后算获胜一次
		turn++
	}
	return win
}
