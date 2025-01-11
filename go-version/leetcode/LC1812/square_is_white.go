package LC1812

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/9 下午10:36
 * @FilePath: square_is_white
 * @Description:
 */

func squareIsWhite(s string) bool {
	return (int(s[0])+int(s[1]))&1 == 1
}
