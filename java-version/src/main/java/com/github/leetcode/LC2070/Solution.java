package com.github.leetcode.LC2070;

import java.util.Arrays;

public class Solution {
    public int[] maximumBeauty(int[][] items, int[] queries) {
        //先对price进行排序
        Arrays.sort(items, (a, b) -> a[0] - b[0]);
        //把beauty换成前缀最大值
        for (int i = 0; i < items.length - 1; i++) {
            items[i+1][1] = Math.max(items[i][1], items[i+1][1]);
        }

        //使用二分查找处理查询
        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            ans[i] = query(items, queries[i]);
        }
        return ans;
    }

    //实现查找逻辑
    private int query(int[][] items, int q) {
        int low = 0, high = items.length;
        while (low < high) {
            int mid = low + ((high - low) >> 1);
            if (items[mid][0] > q) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }
        if (low == 0) {
            return 0;
        }
        return items[low-1][1];
    }
}
