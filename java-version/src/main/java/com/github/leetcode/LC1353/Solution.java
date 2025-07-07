package com.github.leetcode.LC1353;

import java.util.Arrays;
import java.util.PriorityQueue;

public class Solution {
    public int maxEvents(int[][] events) {
        // 先找到会议的最大天数
        int maxDay = 0;
        for (int[] event : events) {
            maxDay = Math.max(maxDay, event[1]);
        }

        // 按照会议的开始时间进行排序
        Arrays.sort(events, (a, b) -> a[0] - b[0]);
        // 以结束时间创建小根堆
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        int ans = 0;
        for (int i = 1, j = 0; i <= maxDay; i++) {
            while (j < events.length && events[j][0] <= i) {
                pq.offer(events[j][1]);
                j++;
            }
            while (!pq.isEmpty() && pq.peek() < i) {
                pq.poll();
            }
            if (!pq.isEmpty()) {
                pq.poll();
                ans++;
            }
        }
        return ans;
    }
}
