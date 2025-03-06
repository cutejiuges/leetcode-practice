package com.github.leetcode.LC2588;

import java.util.HashMap;
import java.util.Map;

/*
* 根据题意可以知，由于每次操作中需要从子数组中选择两个不同的数分别减去 2^k，
* 使得子数组中所有元素均变为 0，由此可知对于子数组中所有元素 2^k 出现的次数之和必须是偶数。
* 换一种说法，即对于二进制中第 i 位，则数组中所有元素第 i 位为 1 的数目一定为偶数，
* 则此时满足数组中所有元素第 i 位的异或和一定为 0。
* */
public class Solution {
    public long beautifulSubarrays(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>();
        int mask = 0;
        long ans = 0;
        map.put(0, 1);
        for (int m : nums) {
            mask ^= m;
            ans += map.getOrDefault(mask, 0);
            map.put(mask, map.getOrDefault(mask, 0) + 1);
        }
        return ans;
    }
}
