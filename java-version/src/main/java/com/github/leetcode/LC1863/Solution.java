package com.github.leetcode.LC1863;

public class Solution {
    private int res = 0, n = 0;

    //深度优先搜索，找到所有的子集
    private void dfs(int val, int idx, int[] nums) {
        if (idx >= this.n) { //递归终止条件
            this.res += val;
            return;
        }
        // 选择当前数字
        dfs(val ^ nums[idx], idx + 1, nums);
        // 跳过当前数字
        dfs(val, idx+1, nums);
    }

    public int subsetXORSum(int[] nums) {
        this.res = 0;
        this.n = nums.length;
        dfs(0, 0, nums);
        return this.res;
    }
}
