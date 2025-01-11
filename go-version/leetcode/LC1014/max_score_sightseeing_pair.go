package LC1014

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/23 下午11:14
 * @FilePath: max_score_sightseeing_pair
 * @Description:
 */

//func maxScoreSightseeingPair(values []int) int {
//	ans := 0
//	for i := 0; i < len(values); i++ {
//		for j := i + 1; j < len(values); j++ {
//			ans = max(ans, values[i]+values[j]+i-j)
//		}
//	}
//	return ans
//}

func maxScoreSightseeingPair(values []int) int {
	//合并同类项，边遍历边维护
	ans, mx := 0, values[0]
	for i := 1; i < len(values); i++ {
		ans = max(ans, mx+values[i]-i)
		mx = max(mx, values[i]+i)
	}
	return ans
}
