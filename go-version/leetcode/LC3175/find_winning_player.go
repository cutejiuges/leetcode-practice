package LC3175

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/24 下午11:05
 * @FilePath: find_winning_player
 * @Description:
 */

func findWinningPlayer(skills []int, k int) int {
	idx, mx := 0, skills[0]
	win := 0

	for i := 1; i < len(skills) && win < k; i++ {
		if skills[i] > mx {
			mx = skills[i]
			idx = i
			win = 1
		} else {
			win++
		}
	}
	return idx
}
