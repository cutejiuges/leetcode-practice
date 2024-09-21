package LC997

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/22 上午12:17
 * @FilePath: find_judge
 * @Description: 出度入度考察，统计计数即可
 */

func findJudge(n int, trust [][]int) int {
	inDegree, outDegree := make([]int, n+1), make([]int, n+1)
	for _, t := range trust {
		outDegree[t[0]]++
		inDegree[t[1]]++
	}

	for i := 1; i <= n; i++ {
		if outDegree[i] == 0 && inDegree[i] == n-1 {
			return i
		}
	}
	return -1
}
