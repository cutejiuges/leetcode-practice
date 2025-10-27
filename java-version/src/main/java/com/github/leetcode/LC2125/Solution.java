package com.github.leetcode.LC2125;

/**
 * 遍历行进行枚举相乘即可
 *
 * @author cutejiuge
 * @since 2025/10/27 下午11:20
 */
public class Solution {
    public int numberOfBeams(String[] bank) {
        int ans = 0, pre = 0;
        for (String ss : bank) {
            int cnt = 0;
            for (int i = 0; i < ss.length(); i++) {
                cnt += ss.charAt(i) - '0';
            }
            if (cnt > 0) {
                ans += cnt * pre;
                pre = cnt;
            }
        }
        return ans;
    }
}
