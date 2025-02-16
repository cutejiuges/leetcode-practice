package com.github.leetcode.LC1299;

public class Solution {
    public int[] replaceElements(int[] arr) {
        int len = arr.length;
        int[] ans = new int[len];
        int mx = -1;
        for(int i = len-1; i>= 0; i--) {
            ans[i] = mx;
            mx = Math.max(mx, arr[i]);
        }
        return ans;
    }
}
