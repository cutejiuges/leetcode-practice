package LC1450

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/1 上午9:12
 * @FilePath: busy_student
 * @Description:
 */

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ans := 0
	for i := range startTime {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			ans++
		}
	}
	return ans
}
