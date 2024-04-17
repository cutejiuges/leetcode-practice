/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-31 10:13:56
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-31 10:53:00
 * @FilePath: /leetcode-practice/leetcode/LC331/is_valid_serialization.go
 * @Description: 二叉树的结构特性问题
 */
package lc331

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}
	inDegree, outDegree := 0, 0
	for i, length := 0, len(preorder); i < length; {
		if i == 0 {
			if preorder[i] == '#' {
				return false
			}
			outDegree = 2
			for i < length && preorder[i] != ',' {
				i++
			}
			continue
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			inDegree++
			i++
		} else {
			inDegree++
			outDegree += 2
			for i < length && preorder[i] != ',' {
				i++
			}
		}
		if i != length && inDegree >= outDegree {
			return false
		}
	}
	return inDegree == outDegree
}
