package math_util

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/16 下午5:38
 * @FilePath: express
 * @Description: 表达式运算
 */

func ThirdExpr(expr bool, a, b any) any {
	if expr {
		return a
	}
	return b
}
