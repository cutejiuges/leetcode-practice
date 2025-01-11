package LC2376

import "strconv"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/20 下午11:25
 * @FilePath: count_special_numbers
 * @Description:
 */

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	type pair struct {
		idx  int
		mask int
	}
	cache := make(map[pair]int)

	var f func(i int, isLimit bool, isNum bool, mask int) int

	f = func(i int, isLimit bool, isNum bool, mask int) (ret int) {
		if i == m {
			if isNum {
				return 1
			} else {
				return 0
			}
		}
		if p, ok := cache[pair{i, mask}]; ok && isNum && !isLimit {
			return p
		}
		var do, up int
		if !isNum {
			ret = f(i+1, false, isNum, mask)
		}
		if !isNum {
			do = 1
		} else {
			do = 0
		}
		if isLimit {
			v, _ := strconv.Atoi(string(s[i]))
			up = v
		} else {
			up = 9
		}
		for k := do; k <= up; k++ {
			if (mask>>(k+1))&1 == 0 {
				ret += f(i+1, isLimit && k == up, true, mask|(1<<(k+1)))
			}
		}
		if isNum && !isLimit {
			cache[pair{i, mask}] = ret
		}
		return
	}

	return f(0, true, false, 0)
}
