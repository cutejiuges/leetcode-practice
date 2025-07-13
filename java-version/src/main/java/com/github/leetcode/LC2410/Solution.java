package com.github.leetcode.LC2410;

import java.util.Arrays;

public class Solution {
    public int matchPlayersAndTrainers(int[] players, int[] trainers) {
        // 先对两个数组进行排序，之后进行遍历比较即可
        Arrays.sort(players);
        Arrays.sort(trainers);
        int m = players.length, n = trainers.length;
        int i = 0, j = 0;
        int ans = 0;
        while (i < m && j < n) {
            if (players[i] <= trainers[j]) {
                ans++;
                i++;
            }
            j++;
        }
        return ans;
    }
}
