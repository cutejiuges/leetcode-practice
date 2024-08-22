package LC3133

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/22 下午11:14
 * @FilePath: min_end
 * @Description:
 */

func minEnd(n, x int) int64 {
	n--
	i, j := 0, 0
	for n>>j > 0 {
		if x>>i&1 == 0 {
			x |= n >> j & 1 << i
			j++
		}
		i++
	}
	return int64(x)
}
