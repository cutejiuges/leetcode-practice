package com.github.datastruct;

/**
 * 定义二叉树结构
 *
 * @author cutejiuge
 * @since 2026/1/7 下午11:38
 */
// Definition for a binary tree node.
public class TreeNode {
    public int val;
    public TreeNode left;
    public TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}
