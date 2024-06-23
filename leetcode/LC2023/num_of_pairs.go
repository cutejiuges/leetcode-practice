package LC2023

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/23 下午9:19
 * @FilePath: num_of_pairs
 * @Description: 枚举前缀和后缀
 */

func numOfPairs(nums []string, target string) int {
	cnt := make(map[string]int)
	ans := 0
	for _, num := range nums {
		cnt[num]++
	}

	for i := 0; i < len(target); i++ {
		prefix, suffix := target[:i], target[i:]
		ans += cnt[prefix] * cnt[suffix]
		if prefix == suffix {
			ans -= cnt[suffix]
		}
	}
	return ans
}
