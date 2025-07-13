package LC2410

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/13 下午9:22
 * @FilePath: match_players_and_trainers
 * @Description:
 */

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	m, n := len(players), len(trainers)
	var i, j, ans int
	for i < m && j < n {
		if players[i] <= trainers[j] {
			ans++
			i++
		}
		j++
	}
	return ans
}
