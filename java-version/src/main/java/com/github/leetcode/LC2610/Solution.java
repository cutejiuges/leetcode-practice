package com.github.leetcode.LC2610;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Solution {
    public List<List<Integer>> findMatrix(int[] nums) {
        // 使用hashMap统计每个数字出现的次数
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            map.merge(num, 1, Integer::sum);
        }
        // 初始化结果数组
        List<List<Integer>> res = new ArrayList<>();

        // 遍历map，填充每一行
        while (!map.isEmpty()) {
            List<Integer> cur = new ArrayList<>(map.keySet());
            res.add(cur);
            for (int m : cur) {
                map.put(m, map.get(m) - 1);
                if (map.get(m) == 0) {
                    map.remove(m);
                }
            }
        }
        return res;
    }
}
