package com.github.leetcode.LC1552;

import java.util.Arrays;

public class Solution {
    public int maxDistance(int[] position, int m) {
        Arrays.sort(position);
        int low = 1, high = position[position.length-1];
        while(low < high) {
            int mid = (low+high+1)>>1;
            if(check(position, m, mid)) {
                low = mid;
            } else {
                high = mid-1;
            }
        }
        return low;
    }

    private boolean check(int[] position, int m, int mid) {
        int pre = position[0], cnt = 1;
        for(int cur : position) {
            if(cur - pre >= mid) {
                cnt++;
                pre = cur;
            }
        }
        return cnt >= m;
    }
}
