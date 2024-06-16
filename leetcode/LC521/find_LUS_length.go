package LC521

import "cutejiuge/leetcode-practice/math_util"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/16 下午5:36
 * @FilePath: find_LUS_length
 * @Description:
 */

func findLUSlength(a, b string) int {
	ans := math_util.ThirdExpr(a == b, -1, max(len(a), len(b)))
	return ans.(int)
}
