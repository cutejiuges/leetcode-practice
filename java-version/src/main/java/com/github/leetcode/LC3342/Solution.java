package com.github.leetcode.LC3342;

import java.util.Arrays;
import java.util.PriorityQueue;

public class Solution {
    // 标记上下左右4个方向
    private final int[][] dirs = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int minTimeToReach(int[][] moveTime) {
        // 创建数组存储到达每一个位置的最小距离
        int row = moveTime.length, col = moveTime[0].length;
        int[][] distance = new int[row][col];
        for (int i = 0; i < row; i++) {
            Arrays.fill(distance[i], Integer.MAX_VALUE);
        }
        // 起点到自己的距离是0
        distance[0][0] = 0;

        // 构建一个最小堆，使用dijkstra算法
        PriorityQueue<int[]> pq = new PriorityQueue<>((a,b) -> a[0] - b[0]);
        pq.offer(new int[]{0, 0, 0}); //起点入队，到自己的距离是0
        while (!pq.isEmpty()) {
            int[] poll = pq.poll();
            int curDist = poll[0], curRow = poll[1], curCol = poll[2];
            // 如果已经到了终点，可以直接返回最小距离
            if (curRow == row - 1 && curCol == col - 1) {
                return curDist;
            }
            // 如果其他位置的更新已经把该位置的值变得更小，则不用再次更新
            if (curDist > distance[curRow][curCol]) {
                continue;
            }
            // 尝试向4个方向各走一步
            for (int[] dir : dirs) {
                int newRow = curRow + dir[0], newCol = curCol + dir[1];
                // 如果该位置不合法，就跳过
                if (newRow < 0 || newRow >= row || newCol < 0 || newCol >= col) {
                    continue;
                }
                // 根据位置来判断走过当前格子需要的消耗
                int time = ((newRow + newCol) & 1) == 0 ? 2 : 1;
                // 走到新的位置上,计算新的距离
                int newDist = Math.max(curDist, moveTime[newRow][newCol]) + time;
                // 判断这个距离是否要比已更新的最小距离小，更小则再次更新
                if (newDist < distance[newRow][newCol]) {
                    distance[newRow][newCol] = newDist;
                    pq.offer(new int[]{newDist, newRow, newCol});
                }
            }
        }
        return -1;
    }
}
