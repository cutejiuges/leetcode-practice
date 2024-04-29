package LC2798

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/30 上午12:12
 * @FilePath: number_of_employees_who_met_target
 * @Description: 简单题一次遍历
 */

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	ans := 0
	for i := range hours {
		if hours[i] >= target {
			ans++
		}
	}
	return ans
}
