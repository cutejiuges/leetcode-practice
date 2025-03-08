package com.github.leetcode.LC2234;

import java.util.Arrays;

public class Solution {
    public long maximumBeauty(int[] flowers, long newFlowers, int target, int full, int partial) {
        int n = flowers.length;

        //全部种满能剩下的数量
        long leftFlowers = newFlowers - (long) target * n;
        for (int i = 0; i < n; i++) {
            flowers[i] = Math.min(target, flowers[i]); //已经满足条件的不种了，当成target处理即可
            leftFlowers += flowers[i]; //把原来的数量补上来，因为已经有的数量不用种
        }

        //如果所有数量都是满的，不需要种，全部都是完善的
        if (leftFlowers == newFlowers) {
            return (long) full * n;
        }

        //如果可以全部种满的话，考虑两种情况，其一是全部种满，其二是第一个少种一个，其他全部种满，取其最大值
        if (leftFlowers >= 0) {
            return Math.max((long) full * n, ((long) full * (n - 1) + (long) partial * (target - 1)));
        }

        //如果不能种满，需要在能种满和种不满之间权衡
        Arrays.sort(flowers);
        long ans = 0, preSum = 0;
        int j = 0;
        //枚举后缀能种满的区间，从1开始，因为0上面已经讨论过了
        for (int i = 1; i <= n; i++) {
            // 撤销在i-1种下的花，不让其满足target
            leftFlowers += target - flowers[i - 1];
            if (leftFlowers < 0) { //剩下的花其实还是不够种的，需要继续撤销
                continue;
            }

            //从0~j位置种花，都能种flowers[j]这么多
            while (j < i && (long) flowers[j] * j <= preSum + leftFlowers) {
                preSum += flowers[j];
                j++;
            }

            //计算当前情况下的总美丽值
            long avg = (leftFlowers + preSum)  / j;
            long beautyValue = avg * partial + (long) full * (n-i);
            ans = Math.max(ans, beautyValue);
        }
        return ans;
    }
}
