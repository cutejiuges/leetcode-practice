package com.github.leetcode.LC1331;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Solution {
    public int[] arrayRankTransform(int[] arr) {
        int[] copy = Arrays.copyOf(arr, arr.length);
        Arrays.sort(copy);
        Map<Integer, Integer> idx = new HashMap<>();
        int sameCnt = 0;
        for (int i = 0; i < copy.length; i++) {
            if (i > 0 && copy[i] == copy[i - 1]) {
                sameCnt++;
            }
            if (!idx.containsKey(copy[i])) {
                idx.put(copy[i], i + 1 - sameCnt);
            }
        }
        int[] ans = new int[arr.length];
        for (int i = 0; i < arr.length; i++) {
            ans[i] = idx.get(arr[i]);
        }
        return ans;
    }
}
