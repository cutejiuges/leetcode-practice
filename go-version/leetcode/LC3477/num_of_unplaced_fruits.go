package LC3477

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/5 下午11:08
 * @FilePath: num_of_unplaced_fruits
 * @Description: 二层循环双重遍历
 */

func numOfUnplacedFruits(fruits, baskets []int) int {
	ans, basketLength := len(fruits), len(baskets)
	flag := make([]bool, basketLength)
	for _, fruit := range fruits {
		for i, v := range baskets {
			if !flag[i] && fruit <= v {
				flag[i] = true
				ans--
				break
			}
		}
	}
	return ans
}
