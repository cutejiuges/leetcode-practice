package com.github.leetcode.LC2141;

/**
 * 二分查找
 *
 * @author cutejiuge
 * @since 2025/12/1 上午8:51
 */
public class Solution {
    /**
     * 如果假设可以运行x分钟，那么每一块电池的时间就是min(battery, x)
     * 对于n台电脑，总共的电量是sum(batteries)，需要的电量是n*x
     * 可以对时间x进行二分，判断能否满足sum(batteries)的条件
     */
    public long maxRunTime(int n, int[] batteries) {
        // 先求出所有电池的电量总和
        long total = 0;
        for (int battery : batteries) {
            total += battery;
        }
        // 二分的左右端点
        long low = 0, high = total / n + 1; // 0是一定满足的，均值+1一定满足不了
        while (low +1 < high) {
            long mid = low + (high - low) / 2;
            long temp = 0;
            for (int battery : batteries) {
                temp += Math.min(mid, battery);
            }
            if (mid * n > temp) {
                high = mid;
            } else {
                low = mid;
            }
        }
        return low;
    }
}
