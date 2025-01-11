package LC2391

import "strings"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/11 上午12:12
 * @FilePath: garbage_collection
 * @Description: 垃圾回收，多次遍历
 */

// 先不管有没有垃圾，全部让垃圾车跑过去并回收，最后逆向遍历一次把时间还回来
func garbageCollection(garbage []string, travel []int) int {
	//先认为每个房子都有三种垃圾，三种车都需要跑到最后回收
	ans := 0
	//记入路程上的时间
	for i := range travel {
		ans += travel[i] * 3
	}
	//记入回收垃圾的时间
	for i := range garbage {
		ans += len(garbage[i])
	}
	n := len(garbage)
	//从最后一个房子开始检查，是否有某种垃圾
	for _, ch := range []byte("MPG") {
		for i := n - 1; i > 0 && strings.IndexByte(garbage[i], ch) < 0; i-- {
			//如果没有的话，不需要回收这种垃圾，把多用的时间还回来
			ans -= travel[i-1]
		}
	}
	return ans
}
