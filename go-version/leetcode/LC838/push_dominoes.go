package LC838

import "strings"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/2 上午12:55
 * @FilePath: push_dominoes
 * @Description: 按照规则进行模拟即可
 */

func pushDominoes(dominoes string) string {
	//1. 在原始数组左侧拼一个L，右侧拼一个R方便处理且不影响结果
	//2. 原始的L、R不会发生变化
	//3. 两个L之间的全部变成L，两个R之间的全部变成R
	//4. 左L和右R之间的全部不变，左R和右L之间的向中间靠
	myDominoes := "L" + dominoes + "R"
	left, right := 0, 1
	var ans strings.Builder
	for right < len(myDominoes) {
		for right < len(myDominoes) && myDominoes[right] == '.' {
			right++
		}
		leftChar, rightChar := myDominoes[left], myDominoes[right]
		if leftChar == rightChar {
			for i := left + 1; i < right; i++ {
				ans.WriteByte(leftChar)
			}
		} else if leftChar == 'R' {
			halfLen := (right - left - 1) >> 1
			for i := 0; i < halfLen; i++ {
				ans.WriteByte('R')
			}
			if (right-left)&1 == 0 {
				ans.WriteByte('.')
			}
			for i := 0; i < halfLen; i++ {
				ans.WriteByte('L')
			}
		} else {
			for i := left + 1; i < right; i++ {
				ans.WriteByte('.')
			}
		}
		ans.WriteByte(rightChar)
		left = right
		right++
	}
	return ans.String()[:ans.Len()-1]
}
