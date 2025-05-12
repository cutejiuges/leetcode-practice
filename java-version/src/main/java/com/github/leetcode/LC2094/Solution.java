package com.github.leetcode.LC2094;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public int[] findEvenNumbers(int[] digits) {
        int[] cnt = new int[10];
        for (int digit : digits) {
            cnt[digit]++;
        }

        List<Integer> ans = new ArrayList<>();
        for (int i = 100; i <= 998; i+=2) {
            if (judge(i, cnt)) {
                ans.add(i);
            }
        }
        return ans.stream().mapToInt(i -> i).toArray();
    }

    private boolean judge(int num, int[] cnt) {
        int[] c = new int[10];
        for (int x = num; x > 0; x /= 10) {
            int cur = x % 10;
            c[cur]++;
            if (c[cur] > cnt[cur]) {
                return false;
            }
        }
        return true;
    }
}
