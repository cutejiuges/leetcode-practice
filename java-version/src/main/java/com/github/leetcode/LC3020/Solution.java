package com.github.leetcode.LC3020;

import java.util.HashMap;
import java.util.Map;

/**
 * LC3020 - 子集中元素的最大数量
 *
 * @author cutejiuge
 * @since 2026/6/27
 */
public class Solution {
    public int maximumLength(int[] nums) {
        Map<Long, Integer> cnt = new HashMap<>();
        for (int num : nums) {
            cnt.merge((long) num, 1, Integer::sum);
        }

        // ans至少是1的数量，向下取奇数
        int oneCnt = cnt.getOrDefault(1L, 0);
        int ans = oneCnt;
        if ((oneCnt & 1) == 0) {
            ans--;
        }

        cnt.remove(1L);

        for (long num : cnt.keySet()) {
            int res = 0;
            long x = num;
            while (cnt.getOrDefault(x, 0) > 1) {
                res += 2;
                x *= x;
            }
            if (cnt.containsKey(x)) {
                ans = Math.max(ans, res + 1);
            } else {
                ans = Math.max(ans, res - 1);
            }
        }
        return ans;
    }
}
