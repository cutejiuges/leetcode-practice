package com.github.leetcode.LC3349;

import java.util.List;

/**
 * 顺序遍历模拟
 *
 * @author cutejiuge
 * @since 2025/10/14 下午10:21
 */
public class Solution {
    public boolean hasIncreasingSubarrays(List<Integer> nums, int k) {
        int n = nums.size();
        for(int i = 0, j, pre = 0, cur; i < n; i++){
            j = i;
            while(i + 1 < n && nums.get(i + 1) > nums.get(i)){
                i++;
            }
            // 连续严格递增 j~i
            cur = i - j + 1;
            // 两种可能性:一个序列>=2k || 前后两个序列均>=k
            if(cur >= k * 2 || (pre >= k && cur >= k)){
                return true;
            }
            pre = cur;
        }
        return false;
    }
}
