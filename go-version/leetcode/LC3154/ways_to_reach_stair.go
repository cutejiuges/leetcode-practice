package LC3154

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/20 下午10:44
 * @FilePath: ways_to_reach_stair
 * @Description: 直接使用记忆化搜索
 */

// 定义一个结构用于追踪递归过程中的链路信息
type traceInfo struct {
	i       int  //记录当前位于第i个阶梯
	j       int  //记录当前已经进行了j次操作二
	preDown bool //记录上一次操作是否是操作一
}

func waysToReachStair(k int) int {
	memo := map[traceInfo]int{}
	var dfs func(int, int, bool) int

	dfs = func(i int, j int, preDown bool) int {
		if i > k+1 { //本次操作无法到达k，直接返回0
			return 0
		}
		cur := traceInfo{i, j, preDown}
		if v, ok := memo[cur]; ok {
			//如果之前已经计算过，记忆化搜索
			return v
		}
		res := dfs(i+1<<j, j+1, false) //进行一次操作二
		if !preDown && i >= 1 {
			//如果上一次不是操作一，且当前能往下走，执行一次操作一
			res += dfs(i-1, j, true)
		}

		//如果执行完操作之后到了第k个台阶，方案数增加
		if i == k {
			res++
		}
		memo[cur] = res
		return res
	}

	return dfs(1, 0, false)
}
