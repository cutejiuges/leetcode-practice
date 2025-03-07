package com.github.leetcode.LC2597;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    private int ans, k;
    private Map<Integer, Integer> map;
    private int[] nums;

    public int beautifulSubsets(int[] nums, int k) {
        this.ans = -1; //去掉回溯中产生的空集
        this.k = k;
        this.nums = nums;
        this.map = new HashMap<>();
        backtrack(0);
        return this.ans;
    }

    private void backtrack(int idx) {
        if (idx == this.nums.length) {
            this.ans++;
            return;
        }
        //不选择当前输入，直接跳过
        backtrack(idx+1);
        //针对当前输入，看能否选择
        int cur = this.nums[idx];
        if (this.map.getOrDefault(cur - this.k, 0) == 0 && this.map.getOrDefault(cur + this.k, 0) == 0) { //当前元素可以选择
            this.map.put(cur, map.getOrDefault(cur, 0) + 1);
            backtrack(idx+1);
            //恢复现场
            this.map.put(cur, map.getOrDefault(cur, 0) - 1);
        }
    }
}
