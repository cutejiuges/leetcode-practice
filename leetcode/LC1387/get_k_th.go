package LC1387

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/22 下午9:02
 * @FilePath: get_k_th
 * @Description: 记忆化搜索
 */

func getKth(lo, hi, k int) int {
	weight := make(map[int]int)
	val := make([]int, 0)
	for i := lo; i <= hi; i++ {
		val = append(val, i)
	}
	sort.Slice(val, func(i, j int) bool {
		if calWeight(val[i], weight) == calWeight(val[j], weight) {
			return val[i] < val[j]
		}
		return calWeight(val[i], weight) < calWeight(val[j], weight)
	})
	return val[k-1]
}

func calWeight(x int, weight map[int]int) int {
	if val, ok := weight[x]; ok {
		return val
	}
	if x == 1 {
		weight[1] = 0
		return 0
	}
	if x&1 == 1 {
		weight[x] = calWeight(3*x+1, weight) + 1
	} else {
		weight[x] = calWeight(x/2, weight) + 1
	}
	return weight[x]
}
