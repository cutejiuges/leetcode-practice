package LC2928

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/1 下午7:06
 * @FilePath: distribute_candies
 * @Description:
 */

func distributeCandies(n int, limit int) int {
	ans := 0
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			if i+j > n {
				break
			}
			if n-(i+j) <= limit {
				ans++
			}
		}
	}
	return ans
}
