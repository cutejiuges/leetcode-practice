package LC2244

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/14 下午11:21
 * @FilePath: minimum_rounds
 * @Description: 统计次数，遍历计算，主要是向上取整的逻辑
 */

func minimumRounds(tasks []int) int {
	//统计每个任务难度出现的频率
	cnt := make(map[int]int)
	for i := range tasks {
		cnt[tasks[i]]++
	}
	ans := 0
	for _, val := range cnt {
		//如果只出现一次，无法拆分为2x+3y
		if val == 1 {
			return -1
		}
		//如果出现了3的倍数，直接除三加上
		if val%3 == 0 {
			ans += val / 3
		} else {
			//如果余2,前面n-1次执行3,最后一次执行2，(3*n+2)/3 = n + ceil(2/3) = n+1
			//如果余1,前面n-2次执行3,最后2次执行2, (3*n+1)/3 = n + ceil(1/3) = n+1
			ans += val/3 + 1
		}
	}
	return ans
}
