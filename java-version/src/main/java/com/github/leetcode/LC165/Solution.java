package com.github.leetcode.LC165;

/**
 * 按字符串切割后，转换数字比较即可
 *
 * @author cutejiuge
 * @since 2025/9/23 下午11:07
 */
public class Solution {
    public int compareVersion(String version1, String version2) {
        String[] v1 = version1.split("\\."), v2 = version2.split("\\.");
        int len1 = v1.length, len2 = v2.length;
        for (int idx = 0; idx < len1 || idx < len2; idx++) {
            int a = idx < len1 ? parseInt(v1[idx]) : 0;
            int b = idx < len2 ? parseInt(v2[idx]) : 0;
            if (a < b) {
                return -1;
            } else if (a > b) {
                return 1;
            }
        }
        return 0;
    }

    private int parseInt(String str) {
        int ans = 0;
        boolean leadZero = true;
        for (int i = 0; i < str.length(); i++) {
            char c = str.charAt(i);
            if (c == '0' && leadZero) {
                continue;
            }else if (c != '0') {
                leadZero = false;
            }
            ans *= 10;
            ans += (c - '0');
        }
        return ans;
    }
}
