package com.github.leetcode.LC1935;

/**
 * 遍历模拟，可以用二进制压缩降低空间占用
 *
 * @author cutejiuge
 * @since 2025/9/15 下午11:21
 */
public class Solution {
    public int canBeTypedWords(String text, String brokenLetters) {
        char[] chars = text.toCharArray(), brokenChars = brokenLetters.toCharArray();
        int mask = 0;
        for (char ch : brokenChars) {
            mask |= (1 << (ch - 'a'));
        }
        int ans = 0, ok = 1;
        for (char ch : chars) {
            if (ch == ' ') {
                ans += ok;
                ok = 1;
            }
            if ((mask & (1 << (ch - 'a'))) != 0) {
                ok = 0;
            }
        }
        ans += ok;
        return ans;
    }
}
