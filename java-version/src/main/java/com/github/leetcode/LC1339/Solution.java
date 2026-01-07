package com.github.leetcode.LC1339;

import com.github.datastruct.TreeNode;

/**
 * 使用两次dfs，分别计算全部的节点和以及当前的节点和
 *
 * @author cutejiuge
 * @since 2026/1/7 下午11:41
 */
public class Solution {
    private static final int MOD = 1_000_000_007;
    private long ans = 0;

    public int maxProduct(TreeNode root) {
        // 先计算当前全部节点和
        long total = dfsTotal(root);
        // 使用全树和再次进行遍历计算，更新结果
        dfsCurrent(root, total);
        // 返回已更新的结果
        return (int) (ans % MOD);
    }

    // 使用dfs计算当前二叉树全部节点的和
    private long dfsTotal(TreeNode root) {
        if (root == null) {
            return 0;
        }
        return root.val + dfsTotal(root.left) + dfsTotal(root.right);
    }

    // 使用dfs计算当前节点的子树和，并更新ans
    private long dfsCurrent(TreeNode root, long total) {
        if (root == null) {
            return 0;
        }
        long cur = root.val + dfsCurrent(root.left, total) + dfsCurrent(root.right, total);
        ans = Math.max(ans, cur * (total - cur));
        return cur;
    }
}
