package LC2079

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/8 上午12:09
 * @FilePath: watering_plants
 * @Description: 植物浇水，直接模拟即可
 */

func wateringPlants(plants []int, capacity int) int {
	ans := 0
	water := capacity
	for i, n := 0, len(plants); i < n; i++ {
		if water >= plants[i] {
			//如果水是够的，当前的水量减掉当前植物的需水量，并前进一步
			water -= plants[i]
			ans++
		} else {
			//如果水不够，返回起点打水，再走过来浇水
			//水量是容量减掉当前植物的需水量，步数是折返再加上当前一步
			ans += i*2 + 1
			water = capacity - plants[i]
		}
	}
	return ans
}
