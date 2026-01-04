package com.github.leetcode.LC1390;

/**
 * 遍历枚举
 *
 * @author cutejiuge
 * @since 2026/1/4 下午9:26
 */
public class Solution {
    public int sumFourDivisors(int[] nums) {
        int ans = 0;
        for (int num : nums) {
            int factorSum = 0, factorCnt = 0;
            for (int x = 1; x * x <= num; x++) {
                if (num % x == 0) {
                    factorSum += x;
                    factorCnt++;
                    if (x * x != num) {
                        factorSum += (num / x);
                        factorCnt++;
                    }
                }
            }
            if (factorCnt == 4) {
                ans += factorSum;
            }
        }
        return ans;
    }
}
