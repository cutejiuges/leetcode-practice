package LC2438

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/11 下午10:46
 * @FilePath: product_queries
 * @Description:
 */

func productQueries(n int, queries [][]int) []int {
	bins := constructBinArray(n)
	return processQueries(bins, queries)
}

func constructBinArray(n int) []int {
	var bins []int
	base := 1
	for n > 0 {
		if n&1 == 1 {
			bins = append(bins, base)
		}
		base <<= 1
		n >>= 1
	}
	return bins
}

func processQueries(bins []int, queries [][]int) []int {
	m := len(bins)
	mod := int(1e9 + 7)
	results := make([][]int, m)
	for i := range results {
		results[i] = make([]int, m)
		cur := 1
		for j := i; j < m; j++ {
			cur = cur * bins[j] % mod
			results[i][j] = cur
		}
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = results[query[0]][query[1]]
	}
	return ans
}
