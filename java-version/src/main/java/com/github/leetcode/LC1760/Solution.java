package com.github.leetcode.LC1760;

public class Solution {
    public int minimumSize(int[] nums, int maxOperations) {
        int mx = 0;
        for (int num: nums) {
            mx = Math.max(mx, num);
        }
        int low = 0, high = mx;
        while(low+1 < high) {
            int mid = (low+high) >> 1;
            if(check(nums, maxOperations, mid)) {
                high = mid;
            } else {
                low = mid;
            }
        }
        return high;
    }

    private boolean check(int[] nums, int maxOperations, int mid) {
        int cnt = 0;
        for (int num: nums) {
            cnt += (num-1) / mid;
        }
        return cnt <= maxOperations;
    }
}
