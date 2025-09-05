package com.github.leetcode.LC2749;

/**
 * 二进制位数判定
 *
 * @author cutejiuge
 * @since 2025/9/5 下午11:26
 */
public class Solution {
    public int makeTheIntegerZero(int num1, int num2) {
        // 尝试枚举步数，假定需要k步，那么转换为num1 = k*(2^i + num2)
        // 所以k*2^i = num1 - k*num2 = x，计算x能否通过k和2^i构成
        for (int k = 1; ; k++) {
            long x = num1 - (long) k * num2;
            if (x < k) {
                return -1;
            }
            if (k >= Long.bitCount(x)) {
                return k;
            }
        }
    }
}
