package com.github.leetcode.LC1590;

import java.util.HashMap;
import java.util.Map;

/**
 * 前缀和 + 哈希
 *
 * @author cutejiuge
 * @since 2025/11/30 下午8:17
 */
public class Solution {
    public int minSubarray(int[] nums, int p) {
        // 记录当前元素和
        long sum = 0;
        for (int num : nums) {
            sum += num;
        }
        int remain = (int) (sum % p);
        if (remain == 0) {
            return 0; // 当前数组不用移除就能整除
        }

        int n = nums.length, ans = nums.length;
        int s = 0;
        Map<Integer, Integer> last = new HashMap<>(); // 记录sum % p最后一次出现的位置
        last.put(s, -1); // 从0开始计数，前缀和从-1开始
        for (int i = 0; i < n; i++) {
            s = (s + nums[i]) % p;
            last.put(s, i); // 更新当前的前缀和 s % p的位置
            int j = last.getOrDefault((s - remain + p) % p, -n); // -n保证i-j为非负数
            ans = Math.min(ans, i - j);
        }
        return ans == n ? -1 : ans;
    }
}
