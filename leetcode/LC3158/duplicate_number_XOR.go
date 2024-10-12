package LC3158

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/12 下午11:30
 * @FilePath: duplicate_number_XOR
 * @Description:
 */

func duplicateNumbersXOR(nums []int) int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	ans := 0
	for num, cnt := range mp {
		if cnt > 1 {
			ans ^= num
		}
	}
	return ans
}
