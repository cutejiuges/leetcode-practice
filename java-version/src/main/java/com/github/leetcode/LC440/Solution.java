package com.github.leetcode.LC440;

public class Solution {
    public int findKthNumber(int n, int k) {
        int cur = 1;
        k--;
        while (k > 0) {
            int step = (int) this.getSteps(cur, n);
            if (step <= k) {
                k -= step;
                cur++;
            } else {
                cur *= 10;
                k--;
            }
        }
        return cur;
    }

    private long getSteps(int cur, long n) {
        long step = 0;
        long first = cur, last = cur;
        while (first <= n) {
            step += (Math.min(last, n) - first) + 1;
            first *= 10;
            last = last * 10 + 9;
        }
        return step;
    }
}
