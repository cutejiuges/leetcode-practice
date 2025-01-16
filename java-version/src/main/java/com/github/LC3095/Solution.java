package com.github.LC3095;

public class Solution {
    public int minimumSubarrayLength(int[] nums, int k) {
        int n = nums.length;
        int[] bits = new int[32];
        int ans = Integer.MAX_VALUE;

        for(int left = 0, right = 0; right < n; right++) {
            for(int i = 0; i < 32; i++) {
                bits[i] += nums[right] >> i & 1;
            }

            while(left <= right && value(bits) >= k) {
                ans = Math.min(ans, right - left + 1);
                for(int i = 0; i < 32; i++) {
                    bits[i] -= nums[left] >> i & 1;
                }
                left++;
            }
        }
        
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }

    private int value(int[] bits) {
        int ans = 0;
        for(int i = 0; i < bits.length; i++) {
            if(bits[i] > 0) {
                ans |= 1 << i;
            }
        }
        return ans;
    }
}
