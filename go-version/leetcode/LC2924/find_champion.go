package LC2924

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/13 下午12:00
 * @FilePath: find_champion
 * @Description: 有向无环图找起点，考虑拓扑排序，但是因为只需要起点，可以弱化进行入度关系查找
 */

// 统计入度
func findChampion(n int, edges [][]int) int {
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	ans, cnt := -1, 0
	for i := range indeg {
		if indeg[i] == 0 {
			ans = i
			cnt++
		}
	}
	if cnt == 1 {
		return ans
	}
	return -1
}

// UnionFindSet 并查集
type UnionFindSet struct {
	Parent []int
}

func initUnionFindSet(n int) *UnionFindSet {
	unionFindSet := &UnionFindSet{
		Parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		unionFindSet.Parent[i] = i
	}
	return unionFindSet
}

func (ufSet *UnionFindSet) find(x int) int {
	if ufSet.Parent[x] == x {
		return x
	}
	ufSet.Parent[x] = ufSet.find(ufSet.Parent[x])
	return ufSet.Parent[x]
}

func (ufSet *UnionFindSet) merge(x, y int) {
	px, py := ufSet.find(x), ufSet.find(y)
	if px == py {
		return
	}
	ufSet.Parent[px] = py
}

func findChampion2(n int, edges [][]int) int {
	ufSet := initUnionFindSet(n)
	for _, e := range edges {
		ufSet.Parent[e[1]] = e[0]
		ufSet.merge(e[0], e[1])
	}
	ans, cnt := -1, 0
	for i := range ufSet.Parent {
		if ufSet.Parent[i] == i {
			ans = i
			cnt++
		}
	}
	if cnt == 1 {
		return ans
	}
	return -1
}
