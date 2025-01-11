package LC3248

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/21 下午11:28
 * @FilePath: final_position_of_snake
 * @Description: 直接遍历就行
 */

func finalPositionOfSnake(n int, commands []string) int {
	ans := 0
	for _, dir := range commands {
		switch dir {
		case "UP":
			ans -= n
		case "DOWN":
			ans += n
		case "LEFT":
			ans--
		case "RIGHT":
			ans++
		}
	}
	return ans
}
