package com.github.leetcode.LC1742;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public int countBalls(int lowLimit, int highLimit) {
        Map<Integer, Integer> map = new HashMap<>();
        int max = 0;
        for(int ballNo = lowLimit; ballNo <= highLimit; ballNo++) {
            int curBall = ballNo, boxNo = 0;
            while(curBall > 0) {
                boxNo += curBall % 10;
                curBall /= 10;
            }
            int cnt = map.getOrDefault(boxNo, 0)+1;
            map.put(boxNo, cnt);
            max = Math.max(cnt, max);
        }
        return max;
    }
}
