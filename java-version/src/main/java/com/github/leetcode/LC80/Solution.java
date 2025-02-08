package com.github.leetcode.LC80;

public class Solution {
    public int removeDuplicates(int[] nums) {
        if(nums == null || nums.length == 0) {
            return 0;
        }
        int curSize = 2, maxSize = nums.length;
        for(int i = 2; i < maxSize; i++) {
            if(nums[i] != nums[curSize - 2]) {
                nums[curSize++] = nums[i];
            }
        }
        return Math.min(curSize, maxSize);
    }
}
