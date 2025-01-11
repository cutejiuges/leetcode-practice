package LC3274

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/3 下午10:27
 * @FilePath: check_two_chessboards
 * @Description: 找规律，街区距离为奇数的位置一定是异色，偶数的位置一定是同色
 */

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	return (int(coordinate1[0]-coordinate2[0])+int(coordinate1[1]-coordinate2[1]))&1 == 0
}
