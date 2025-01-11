package LC3222

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/5 下午11:37
 * @FilePath: losing_player
 * @Description:
 */

func losingPlayer(x int, y int) string {
	return [2]string{"Bob", "Alice"}[min(x, y/4)&1]
}
