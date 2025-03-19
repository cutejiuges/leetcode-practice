package LC2610

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/19 下午11:15
 * @FilePath: find_matrix
 * @Description:
 */

func findMatrix(nums []int) [][]int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	res := make([][]int, 0)
	for len(cnt) > 0 {
		cur := make([]int, 0)
		for k, v := range cnt {
			cur = append(cur, k)
			cnt[k] = v - 1
			if cnt[k] == 0 {
				delete(cnt, k)
			}
		}
		res = append(res, cur)
	}
	return res
}
