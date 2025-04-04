package com.github.leetcode.LC1123;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
}


public class Solution {
    private TreeNode ans;
    private int maxDepth = -1; // 全局最大深度

    public TreeNode lcaDeepestLeaves(TreeNode root) {
        dfs(root, 0);
        return ans;
    }

    private int dfs(TreeNode node, int depth) {
        if (node == null) {
            maxDepth = Math.max(maxDepth, depth); // 维护全局最大深度
            return depth;
        }
        int leftMaxDepth = dfs(node.left, depth + 1); // 左子树最深空节点的深度
        int rightMaxDepth = dfs(node.right, depth + 1); // 右子树最深空节点的深度
        if (leftMaxDepth == rightMaxDepth && leftMaxDepth == maxDepth) { // 最深的空节点左右子树都有
            ans = node;
        }
        return Math.max(leftMaxDepth, rightMaxDepth); // 当前子树最深空节点的深度
    }
}
