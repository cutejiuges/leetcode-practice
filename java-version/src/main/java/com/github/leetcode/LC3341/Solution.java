package com.github.leetcode.LC3341;

import java.util.Arrays;
import java.util.PriorityQueue;

public class Solution {
    // 标记上下左右4个移动方向
    private final int[][] dirs = new int[][]{{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int minTimeToReach(int[][] minTime) {
        int row = minTime.length, col = minTime[0].length;
        int[][] distance = new int[row][col];
        for (int i = 0; i < row; i++) {
            Arrays.fill(distance[i], Integer.MAX_VALUE);
        }
        distance[0][0] = 0; //从起点到起点的距离是0

        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        pq.offer(new int[]{0, 0, 0}); // 起点为0,0且距离为0
        while (!pq.isEmpty()) {
            int[] poll = pq.poll();
            int curDistance = poll[0], curRow = poll[1], curCol = poll[2];
            // 当前位置已经在终点了，返回当前距离
            if (curRow == row - 1 && curCol == col - 1) {
                return curDistance;
            }
            // 如果到达该位置的距离已经在更新过程中变得更小了，则可以不用进行这一次迭代
            if (distance[curRow][curCol] < curDistance) {
                continue;
            }
            for (int[] dir : dirs) {
                int newRow = curRow + dir[0], newCol = curCol + dir[1];
                // 如果向该方向走过去位置不合法，跳过这个位置
                if (newRow < 0 || newRow >= row || newCol < 0 || newCol >= col) {
                    continue;
                }
                // 新位置的距离是当前最小距离和该位置限制距离，加上移动一次的距离
                int newDistance = Math.max(minTime[newRow][newCol], curDistance) + 1;
                // 如果计算距离更小，就执行更新
                if (newDistance < distance[newRow][newCol]) {
                    distance[newRow][newCol] = newDistance;
                    pq.offer(new int[]{newDistance, newRow, newCol});
                }
            }
        }
        return -1;
    }
}
