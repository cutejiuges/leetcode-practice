package com.github.leetcode.LC2327;

import java.util.ArrayDeque;
import java.util.Deque;

/**
 * 使用两个双端队列进行模拟
 *
 * @author cutejiuge
 * @since 2025/9/9 下午11:10
 */
public class Solution {
    public int peopleAwareOfSecret(int n, int delay, int forget) {
        int mod = (int) 1e9 + 7;
        // 两个双端队列，分别表示仅知道秘密和会分享秘密的人
        Deque<int[]> know = new ArrayDeque<>(), share = new ArrayDeque<>();
        // 第1天只有1个人知道秘密
        know.offerLast(new int[]{1,1});
        // 两个整数，分别表示仅知道秘密的人数和会分享秘密的人数
        int knowCnt = 1, shareCnt = 0;

        // 从第2天开始模拟
        for (int i = 2; i <= n; i++) {
            // 延迟的天数到了，know队列中满足条件的可以开始进行传播了
            if (!know.isEmpty() && know.peekFirst()[0] == i - delay) {
                int[] head = know.pollFirst();
                // 仅知道秘密的人数减少，会传播秘密的人数增加
                knowCnt = (knowCnt - head[1] + mod) % mod;
                shareCnt = (shareCnt + head[1]) % mod;
                share.offerLast(head);
            }
            // 忘记的天数到了，share中的元素可以删除
            if (!share.isEmpty() && share.peekFirst()[0] == i - forget) {
                int[] head = share.pollFirst();
                shareCnt = (shareCnt - head[1] + mod) % mod;
            }
            // 开始传播秘密
            if (!share.isEmpty()) {
                // 知道秘密的人数是当前知道秘密的人数加上可以分享秘密的人数
                knowCnt = (knowCnt + shareCnt) % mod;
                // 新加上了一部分仅知道秘密的人数
                know.offerLast(new int[]{i, shareCnt});
            }
        }
        // 最终知道秘密的人数是仅知道秘密的人加上能传播秘密的人
        return (knowCnt + shareCnt) % mod;
    }
}
