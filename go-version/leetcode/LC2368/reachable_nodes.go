/*
* @Author: cutejiuge cutejiuge@163.com
* @Date: 2024-03-02 20:25:10
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-02 20:48:12
* @FilePath: /leetcode-practice/LC2368/reachable_nodes.go
* @Description:leetcode 2368,可以使用深度优先遍历方式解决问题，转换成图论问题，即求跟节点所在的最大联通块的大小，
* dfs从根节点可以遍历到的大小，或者跟节点的并查集都可以，这里用dfs
*/
package lc2368

func reachableNodes(n int, edges [][]int, restricted []int) int {
	//记录已经被禁止的节点，空间换时间
	isRestricted := make(map[int]bool)
	for _, v := range restricted {
		isRestricted[v] = true
	}

	//把边的集合转换成邻接矩阵，空间换时间
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	//通过深度有点搜索的方式，记录从跟节点可以到达的节点数量
	cnt := 0 //记录可达的节点总数
	var dfs func(int, int)
	dfs = func(cur, parent int) {
		cnt++
		for _, next := range graph[cur] {
			if next != parent && !isRestricted[next] {
				dfs(next, cur)
			}
		}
	}
	dfs(0, -1)
	return cnt
}
