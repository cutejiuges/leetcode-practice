package LC3181

import (
	"math/big"
	"sort"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/26 下午6:50
 * @FilePath: max_total_reward
 * @Description:
 */

func maxTotalReward(rewardValues []int) int {
	n := len(rewardValues)
	sort.Ints(rewardValues)
	if n >= 2 && rewardValues[n-2] == rewardValues[n-1]-1 {
		return rewardValues[n-1]*2 - 1
	}

	f0, f1 := big.NewInt(1), big.NewInt(0)
	for _, x := range rewardValues {
		mask, one := big.NewInt(0), big.NewInt(1)
		mask.Sub(mask.Lsh(one, uint(x)), one)
		f1.Lsh(f1.And(f0, mask), uint(x))
		f0.Or(f1, f0)
	}
	return f0.BitLen() - 1
}
