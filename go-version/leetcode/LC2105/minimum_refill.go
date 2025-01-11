package LC2105

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/10 下午11:52
 * @FilePath: minimum_refill
 * @Description: 植物浇水，还是直接模拟
 */

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	ans, n := 0, len(plants)
	posA, posB := 0, n-1
	waterA, waterB := capacityA, capacityB
	for posA < posB {
		if waterA >= plants[posA] {
			waterA -= plants[posA]
		} else {
			ans++
			waterA = capacityA - plants[posA]
		}
		posA++

		if waterB >= plants[posB] {
			waterB -= plants[posB]
		} else {
			ans++
			waterB = capacityB - plants[posB]
		}
		posB--
	}
	if posA == posB {
		if (waterA >= waterB && waterA < plants[posA]) || (waterB > waterA && waterB < plants[posB]) {
			ans++
		}
	}
	return ans
}
