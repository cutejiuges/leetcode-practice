package com.github.leetcode.LC1399;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Solution {
    public int countLargestGroup(int n) {
        Map<Integer, List<Integer>> map = new HashMap<>();
        int maxSize = 0, ans = 0;
        for (int i = 1; i <= n; i++) {
            int sum = this.getDigitSum(i);
            List<Integer> curList = map.getOrDefault(sum, new ArrayList<>());
            curList.add(i);
            map.put(sum, curList);
            if (maxSize < curList.size()) {
                maxSize = curList.size();
                ans = 1;
            } else if (maxSize == curList.size()) {
                ans++;
            }
        }
        return ans;
    }

    private int getDigitSum(int x) {
        int sum = 0, t = x;
        while (t > 0) {
            sum += t % 10;
            t /= 10;
        }
        return sum;
    }
}
