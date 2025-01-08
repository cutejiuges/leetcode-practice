package LC2264

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/8 下午8:56
 * @FilePath: largest_good_integer
 * @Description:
 */

func largestGoodInteger(num string) string {
	n, ans := len(num), ""
	for i := 0; i < n-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			ans = max(ans, num[i:i+3])
		}
	}
	return ans
}
