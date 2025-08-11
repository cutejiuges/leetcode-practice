package com.github.leetcode.LC2438;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public int[] productQueries(int n, int[][] queries) {
        List<Integer> bins = constructBinArray(n);
        return processQueries(bins, queries);
    }

    private List<Integer> constructBinArray(int n) {
        List<Integer> bins = new ArrayList<>();
        int base = 1;
        while (n > 0) {
            if ((n & 1) == 1) {
                bins.add(base);
            }
            base <<= 1;
            n >>= 1;
        }
        return bins;
    }

    private int[] processQueries(List<Integer> bins, int[][] queries) {
        int len = bins.size(), mod = (int) 1e9 + 7;
        int[][] results = new int[len][len];
        for (int i = 0; i < len; i++) {
            long cur = 1;
            for (int j = i; j < len; j++) {
                cur  = (cur % mod) * (bins.get(j) % mod) % mod;
                results[i][j] = (int) cur;
            }
        }
        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            ans[i] = results[queries[i][0]][queries[i][1]];
        }
        return ans;
    }
}
