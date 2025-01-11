package LC3249

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/14 下午11:16
 * @FilePath: count_good_nodes
 * @Description: 深度优先搜索，统计各个节点的子树
 */

func countGoodNodes(edges [][]int) int {
	//构建邻接表
	n := len(edges) + 1
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	ans := 0
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		flag := true     //标记当前节点是否为好节点
		treeSize := 0    //当前节点为根的子树大小
		subTreeSize := 0 // 某个子节点为根的子树大小

		//遍历当前节点的子结点
		for _, child := range graph[node] {
			if child != parent { //邻接表，避开父节点
				size := dfs(child, node) //求出当前子结点的大小
				if subTreeSize == 0 {    //如果之前没算过，就算作第一个
					subTreeSize = size
				} else if subTreeSize != size { //如果当前子结点的大小和前面子结点的大小不一样，则当前节点不是好节点
					flag = false
				}
				treeSize += size //当前节点的数规模，累加子结点树
			}
		}

		if flag { //如果是好节点，数量++
			ans++
		}
		return treeSize + 1 //返回当前节点为根的数大小，需要加上自己
	}

	dfs(0, -1)
	return ans
}
