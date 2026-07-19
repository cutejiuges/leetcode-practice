package com.github.leetcode.LC1081;

public class Solution {
    public String smallestSubsequence(String s) {
        int[] cnt = new int[26];
        for (int i = 0; i < s.length(); i++) {
            cnt[s.charAt(i) - 'a']++;
        }
        boolean[] visited = new boolean[26];
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < s.length(); i++) {
            char ch = s.charAt(i);
            cnt[ch - 'a']--;
            // 避免重复
            if (visited[ch - 'a']) continue;
            // 如果结果的顶部字符比当前字符更大，且待选集后面还有这个字符，就先把这个字符抛出来
            while (sb.length() > 0 && ch < sb.charAt(sb.length() - 1) && cnt[sb.charAt(sb.length() - 1) - 'a'] > 0) {
                visited[sb.charAt(sb.length() - 1) - 'a'] = false;
                sb.deleteCharAt(sb.length() - 1);
            }
            sb.append(ch);
            visited[ch - 'a'] = true;
        }
        return sb.toString();
    }
}
