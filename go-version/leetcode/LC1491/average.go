package LC1491

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午1:08
 * @FilePath: average
 * @Description: 去掉最高分和最低分之后的平均数
 */

func average(salary []int) float64 {
	sum, maxSalary, minSalary, n := float64(salary[0]), salary[0], salary[0], len(salary)
	for i := 1; i < n; i++ {
		if salary[i] < minSalary {
			minSalary = salary[i]
		}
		if salary[i] > maxSalary {
			maxSalary = salary[i]
		}
		sum += float64(salary[i])
	}
	sum -= float64(minSalary + maxSalary)
	return sum / (float64(n - 2))
}
