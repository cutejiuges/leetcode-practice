package com.github.leetcode.LC869;

public class Solution {
    private boolean[] visited;

    public boolean reorderedPowerOf2(int n) {
        char[] chars = String.valueOf(n).toCharArray();
        visited = new boolean[chars.length];
        return backtrack(chars, 0, 0);
    }

    private boolean powOf2(int n) {
        return n > 0 && (n & (n - 1)) == 0;
    }

    private boolean backtrack(char[] nums, int idx, int num) {
        if (idx == nums.length) {
            return powOf2(num);
        }

        for (int i = 0; i < nums.length; i++) {
            // 排除前导0元素
            if (nums[i] == '0' && num == 0) {
                continue;
            }
            // 已经访问过不能再次访问
            if (visited[i]) {
                continue;
            }
            visited[i] = true;
            if (backtrack(nums, idx + 1, num * 10 + nums[i] - '0')) {
                return true;
            }
            visited[i] = false;
        }
        return false;
    }
}
