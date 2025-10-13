package com.github.leetcode.LC2273;

import java.util.ArrayList;
import java.util.List;

/**
 * 顺序遍历判断即可
 *
 * @author cutejiuge
 * @since 2025/10/13 下午11:05
 */
public class Solution {
    public List<String> removeAnagrams(String[] words) {
        List<String> ans = new ArrayList<>();
        if (words == null || words.length == 0) {
            return ans;
        }
        ans.add(words[0]);
        int idx = 0;
        for (int i = 1; i < words.length; i++) {
            if (anagram(ans.get(idx), words[i])) {
                continue;
            }
            ans.add(words[i]);
            idx++;
        }
        return ans;
    }

    private boolean anagram(String a, String b) {
        if (a.length() != b.length()) {
            return false;
        }
        int[] cnt = new int[26];
        for (int i = 0; i < a.length(); i++) {
            cnt[a.charAt(i) - 'a']++;
        }
        for (int i = 0; i < b.length(); i++) {
            cnt[b.charAt(i) - 'a']--;
        }
        for (int i = 0; i < cnt.length; i++) {
            if (cnt[i] != 0) {
                return false;
            }
        }
        return true;
    }
}
