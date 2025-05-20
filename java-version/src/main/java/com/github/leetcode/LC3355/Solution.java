package com.github.leetcode.LC3355;

public class Solution {
    public boolean isZeroArray(int[] nums, int[][] queries) {
        /*使用差分数组进行数据判断*/
        int[] diffArray = new int[nums.length];
        for (int[] query : queries) {
            diffArray[query[0]]++;
            if (query[1] < nums.length - 1) {
                diffArray[query[1] + 1]--;
            }
        }
        int sumDiff = 0;
        for (int i = 0; i < nums.length; i++) {
            sumDiff += diffArray[i];
            if (nums[i] > sumDiff) {
                return false;
            }
        }
        return true;
    }
}
