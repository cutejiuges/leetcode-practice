package LC684

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/27 上午1:47
 * @FilePath: find_redundant_connection
 * @Description: 使用并查集思想来实现
 */

func findRedundantConnection(edges [][]int) []int {
	//初始时,将每一个节点都作为一个连同子域
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}

	//并查集的连通
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			return false
		}
		parent[px] = py
		return true
	}

	//遍历每一个边,进行节点的连通操作
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}
