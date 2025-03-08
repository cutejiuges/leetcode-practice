package LC2234

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/8 下午12:30
 * @FilePath: maximum_beauty
 * @Description: 前缀和验证博弈
 */

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	n := len(flowers)

	//先把每个位置都种满target，看看能剩下多少花
	leftFlower := newFlowers - int64(target*n)
	for i := 0; i < n; i++ { //把已经有的加上去，这部分不用种
		flowers[i] = min(target, flowers[i])
		leftFlower += int64(flowers[i])
	}

	//如果本身全是满的，啥都不用种
	if leftFlower == newFlowers {
		return int64(full * n)
	}

	//如果需要种，但是可以全部种满，有两种方式，其一全部种满，其二留一个剩一朵，别的全种满
	if leftFlower >= 0 {
		return max(int64(partial*(target-1)+full*(n-1)), int64(full*n))
	}

	sort.Ints(flowers)
	//如果种不满，需要枚举能种满的后缀和种不满的前缀，选出最大的
	var ans, preSum int64
	j := 0
	//枚举后缀，从1开始，因为前面已经算过全部种满的情况了
	for i := 1; i <= n; i++ {
		//撤销在i-1位置种的花
		leftFlower += int64(target - flowers[i-1])
		if leftFlower < 0 { //撤销了还是不够种满后面的，需要继续撤销
			continue
		}

		//如果后面种满还有的剩，就往前面种，种到flowers[j]
		for j < i && int64(flowers[j]*j) <= leftFlower+preSum {
			preSum += int64(flowers[j])
			j++
		}
		//计算此时的美丽值
		avg := (preSum + leftFlower) / int64(j)
		beautyValue := int64(full*(n-i)) + int64(partial)*avg
		ans = max(ans, beautyValue)
	}
	return ans
}
