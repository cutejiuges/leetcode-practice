package com.github.leetcode.LC865;

import com.github.datastruct.TreeNode;

/**
 * 递归树高
 *
 * @author cutejiuge
 * @since 2026/1/9 下午11:21
 */
public class Solution {
    private int depth = -1; // 全局的最大树高
    private TreeNode ans = null;

    public TreeNode subtreeWithAllDeepest(TreeNode root) {
        dfs(root, 0);
        return this.ans;
    }

    private int dfs(TreeNode root, int depth) {
        if (root == null) {
            this.depth = Math.max(this.depth, depth); // 更新全局最大树高
            return depth; // 返回当前树高
        }
        int leftDepth = this.dfs(root.left, depth + 1); // 维护左子树的最大树高
        int rightDepth = this.dfs(root.right, depth + 1); // 维护右子树的最大树高
        if (leftDepth == this.depth && rightDepth == this.depth) {
            this.ans = root; // 如果当前节点的左树和右树一样高，且都和最深的树高一致，就更新答案
        }
        // 更新答案之后还需要进行递归，后面可能还会有更深的节点满足条件，此处需要返回当前树高
        return Math.max(leftDepth, rightDepth);
    }
}
