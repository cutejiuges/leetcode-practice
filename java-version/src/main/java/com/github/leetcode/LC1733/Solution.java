package com.github.leetcode.LC1733;

import java.util.HashSet;
import java.util.Set;

/**
 * 贪心算法，统计多少好友不能沟通，掌握最多的语言是什么
 *
 * @author cutejiuge
 * @since 2025/9/10 下午10:59
 */
public class Solution {
    public int minimumTeachings(int n, int[][] languages, int[][] friendships) {
        // 存储不能沟通的人
        Set<Integer> notCommunicate = new HashSet<>();
        // 遍历好友对
        for (int[] friend : friendships) {
            Set<Integer> set = new HashSet<>();  // 存放当前好友对中0会的语言
            boolean flag = false;  // 能否沟通
            for (int lan : languages[friend[0] - 1]) { // 将当前好友对中，0会的语言塞进去
                set.add(lan);
            }
            // 遍历1会的语言，看看是否有都会的
            for (int lan : languages[friend[1] - 1]) {
                if (set.contains(lan)) {
                    flag = true;
                    break;
                }
            }
            // 如果没有共同语言，认为这两个人不能沟通
            if (!flag) {
                notCommunicate.add(friend[0] - 1);
                notCommunicate.add(friend[1] - 1);
            }
        }
        // 不同语言的人中间，最多人会讲得那一种数量
        int maxCnt = 0;
        int[] cnt = new int[n + 1];
        for (int friend : notCommunicate) {
            for (int lan : languages[friend]) {
                cnt[lan]++;
                maxCnt = Math.max(maxCnt, cnt[lan]);
            }
        }
        // 其他的人都学会最多人会的就行了
        return notCommunicate.size() - maxCnt;
    }
}
