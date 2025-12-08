package com.github.leetcode.LC1925;

/**
 * 枚举遍历
 *
 * @author cutejiuge
 * @since 2025/12/8 上午8:15
 */
public class Solution {
    public int countTriples(int n) {
        int ans = 0;
        for (int a = 1; a < n; a++) { // 枚举a
            for (int b = 1; b < a && a * a + b * b <= n * n; b++) { // 枚举满足条件的b
                int c = a * a + b * b;   // 计算得到斜边方
                int rt = (int) Math.sqrt(c); // 计算斜边
                if (rt * rt == c) { // 如果正好满足条件则计入，否则说明不是整数有问题
                    ans++;
                }
            }
        }
        return ans * 2;
    }
}
