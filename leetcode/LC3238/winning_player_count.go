package LC3238

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/23 下午11:23
 * @FilePath: winning_player_count
 * @Description:
 */

func winningPlayerCount(n int, pick [][]int) int {
	colorCnt := make([][11]int, n)
	for _, item := range pick {
		colorCnt[item[0]][item[1]]++
	}

	ans := 0
	for i, cnt := range colorCnt {
		for _, c := range cnt {
			if c > i {
				ans++
				break
			}
		}
	}
	return ans
}
