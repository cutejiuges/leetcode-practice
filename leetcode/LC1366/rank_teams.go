package LC1366

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/29 下午9:53
 * @FilePath: rank_teams
 * @Description:
 */

func rankTeams(votes []string) string {
	//统计每一个字符ch在某个位置i出现的次数
	cnts := make(map[rune][]int)
	for _, ch := range votes[0] {
		cnts[ch] = make([]int, len(votes[0]))
	}

	for _, vote := range votes {
		for i, ch := range vote {
			cnts[ch][i]++
		}
	}

	//定义一个临时结构，标记某个字符，以及字符在每一个位置出现的次数
	type tp struct {
		ch   rune
		rank []int
	}
	//从map里面抽出来，放进一个数组里面
	ranks := make([]tp, 0, len(cnts))
	for k, v := range cnts {
		ranks = append(ranks, tp{k, v})
	}

	//对数组进行排序
	sort.Slice(ranks, func(i, j int) bool {
		//从位置进行遍历，捞出来这个位置最多的，如果这个位置一样就顺延到下一个位置
		for k := 0; k < len(ranks[i].rank); k++ {
			if ranks[i].rank[k] != ranks[j].rank[k] {
				return ranks[i].rank[k] > ranks[j].rank[k]
			}
		}
		//全部一样的话就直接按照字符顺序排序
		return ranks[i].ch < ranks[j].ch
	})

	//按照顺序把字符组合到一起来即可
	ans := make([]rune, 0, len(ranks))
	for _, t := range ranks {
		ans = append(ans, t.ch)
	}
	return string(ans)
}
