package com.github.leetcode.LC3477;

public class Solution {
    public int numOfUnplacedFruits(int[] fruits, int[] baskets) {
        boolean[] flag = new boolean[baskets.length];
        int ans = fruits.length;
        for (int fruit : fruits) {
            for (int i = 0; i < baskets.length; i++) {
                if (!flag[i] && fruit <= baskets[i]) {
                    flag[i] = true;
                    ans--;
                    break;
                }
            }
        }
        return ans;
    }
}
