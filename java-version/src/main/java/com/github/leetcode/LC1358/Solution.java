package com.github.leetcode.LC1358;

public class Solution {
    public int numberOfSubstrings(String s) {
        int ans = 0, left = 0;
        int[] cnt = new int[3];
        int length = s.length();
        for (int i = 0; i < length; i++) {
            cnt[s.charAt(i) - 'a']++;
            while (cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0) {
                cnt[s.charAt(left) - 'a']--;
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
