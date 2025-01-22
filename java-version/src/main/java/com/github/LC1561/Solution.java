package com.github.LC1561;

import java.util.Arrays;

public class Solution {
    public int maxCoins(int[] piles) {
        Arrays.sort(piles);
        int len = piles.length, round = len / 3;
        int ans = 0, idx = len - 2;
        for(int i = 0; i < round; i++) {
            ans += piles[idx];
            idx -= 2;
        }   
        return ans;
    }
}
