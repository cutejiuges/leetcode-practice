package com.github.leetcode.LC2943;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

/**
 * 最长连续序列优化方式
 *
 * @author cutejiuge
 * @since 2026/1/15 下午10:45
 */
public class Solution {
    public int maximizeSquareHoleArea(int n, int m, int[] hBars, int[] vBars) {
        int d = Math.min(this.longestConsecutive(hBars), this.longestConsecutive(vBars)) + 1;
        return d * d;
    }

    private int longestConsecutive(int[] nums) {
        // 数组转换成集合
        Set<Integer> set = new HashSet<>();
        for (int num : nums) {
            set.add(num);
        }

        int ans = 0;
        for (int x : set) {
            if (set.contains(x - 1)) { // 如果x不是序列的起点就直接跳过
                continue;
            }
            // 当x是序列的起点不断寻找下一个数直到序列结束
            int y = x + 1;
            while (set.contains(y)) {
                ++y;
            }
            // 结束之后更新当前最大的序列
            ans = Math.max(ans, y - x);
        }
        return ans;
    }
}
