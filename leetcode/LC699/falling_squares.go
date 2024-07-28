package LC699

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/28 上午8:04
 * @FilePath: falling_squares
 * @Description:
 */

type Interval struct {
	left, right, height int
}

func fallingSquares(positions [][]int) []int {
	fall := make([]Interval, 0)
	ans := make([]int, len(positions))

	for i, pos := range positions {
		process(pos, &fall, i, ans)
	}
	return ans
}

func process(pos []int, fall *[]Interval, ai int, ans []int) {
	//该位置的方块信息
	left, right, height := pos[0], pos[0]+pos[1], pos[1]
	highest := 0

	//遍历已落地的方块信息，判断能开始叠加的初始高度
	for _, interval := range *fall {
		if interval.left >= right || interval.right <= left {
			continue
		}
		//否则把当前位置的高度拿出来比较，取最大的
		highest = max(highest, interval.height)
	}
	//以上计算该位置的方块下落后，底边的高度，此时将该方块加入已落地的列表
	*fall = append(*fall, Interval{left: left, right: right, height: height + highest})

	//如果是首个方块下落，一定是height+highest，后续需要和前面的最大高度取一个最大值
	if ai == 0 {
		ans[ai] = height + highest
	} else {
		ans[ai] = max(height+highest, ans[ai-1])
	}
}
