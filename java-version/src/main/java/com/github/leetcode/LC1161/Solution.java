package com.github.leetcode.LC1161;

import com.github.datastruct.TreeNode;

import java.util.ArrayDeque;
import java.util.Queue;

/**
 * 层序遍历求最大值即可
 *
 * @author cutejiuge
 * @since 2026/1/6 下午8:39
 */
public class Solution {
    public int maxLevelSum(TreeNode root) {
        int ans = 0, layer = 0;
        long maxSum = Long.MIN_VALUE;
        // 使用队列存储每一层的节点
        Queue<TreeNode> queue = new ArrayDeque<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            ++layer;
            int curSize = queue.size();
            long curSum = 0;
            for (int i = 0; i < curSize; ++i) {
                TreeNode node = queue.poll();
                if (node == null) continue;
                curSum += node.val;
                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }
            ans = curSum > maxSum ? layer : ans;
            maxSum = Math.max(curSum, maxSum);
        }
        return ans;
    }
}
