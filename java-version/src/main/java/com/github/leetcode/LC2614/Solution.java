package com.github.leetcode.LC2614;

import java.util.Arrays;

public class Solution {
    private final boolean[] isPrime = new boolean[(int) (4*1e6+1)];

    public int diagonalPrime(int[][] nums) {
        initPrimes();
        int ans = 0, n = nums.length;
        for (int i = 0; i < n; i++) {
            if (this.isPrime[nums[i][i]]) {
                ans = Math.max(ans, nums[i][i]);
            }
            if (this.isPrime[nums[i][n - i - 1]]) {
                ans = Math.max(ans, nums[i][n - i - 1]);
            }
        }
        return ans;
    }

    //艾氏筛进行素数判断
    private void initPrimes() {
        Arrays.fill(this.isPrime, true);
        this.isPrime[0] = this.isPrime[1] = false;
        for (int i = 2; i*i < this.isPrime.length; i++) {
            if (this.isPrime[i]) {
                for (int j = i*i; j < this.isPrime.length; j += i) {
                    this.isPrime[j] = false;
                }
            }
        }
    }
}
