/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-12 00:21:18
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-12 01:10:01
 * @FilePath: /leetcode-practice/data_struct/lc_tree.go
 * @Description: 本文件实现leetcode常用的树结构及常用操作
 */
package datastruct

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(elems []string, index int) *TreeNode {
	n := len(elems)
	if index >= n || elems[index] == "null" {
		return nil
	}
	nodeVal, _ := strconv.Atoi(elems[index])
	root := &TreeNode{
		Val:   nodeVal,
		Left:  createTree(elems, 2*index+1),
		Right: createTree(elems, 2*index+2),
	}
	return root
}

func Construct(elems []string) (root *TreeNode) {
	root = createTree(elems, 0)
	return root
}

func PreOrder(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func MiddleOrder(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func PostOrder(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}
