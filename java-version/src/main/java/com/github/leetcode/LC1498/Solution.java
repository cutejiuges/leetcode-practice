package com.github.leetcode.LC1498;

import java.util.Arrays;

public class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int[] pow2 = new int[100000];
    private static boolean initializes = false;

    private void init() {
        if (initializes) {
            return;
        }
        initializes = true;
        pow2[0] = 1;
        for (int i = 1; i < pow2.length; i++) {
            pow2[i] = pow2[i - 1] * 2 % MOD;
        }
    }

    public int numSubseq(int[] nums, int target) {
        init();
        Arrays.sort(nums);
        long ans = 0;
        int low = 0, high = nums.length - 1;
        while (low <= high) {
            if (nums[low] + nums[high] <= target) {
                ans += pow2[high - low];
                low++;
            } else {
                high--;
            }
        }
        return (int) (ans % MOD);
    }
}
