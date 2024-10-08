package LC1436

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/8 上午8:51
 * @FilePath: dest_city
 * @Description:
 */

// 统计入度，拿出最终出度为0的节点
func destCity(paths [][]string) string {
	outDegree := make(map[string]struct{})
	for _, path := range paths {
		outDegree[path[0]] = struct{}{}
	}

	for _, path := range paths {
		if _, ok := outDegree[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}
