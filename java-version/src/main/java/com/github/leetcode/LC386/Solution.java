package com.github.leetcode.LC386;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public List<Integer> lexicalOrder(int n) {
        List<Integer> ans = new ArrayList<>();
        // 从1开始构造
        int num = 1;
        // 总共构造n个数
        for (int i = 0; i < n; i++) {
            ans.add(num); // 将每次构造的结果加入结果集
            if (num * 10 <= n) { // 如果乘以10还在范围内，优先乘以10保证字典序
                num *= 10;
            } else { // 否则的话需要分情况讨论
                while (num % 10 == 9 || num >= n) { // 如果已经到了范围，或者本轮的禁止已经到了9
                    num /= 10; //回退到上一轮的结果
                }
                num++; //字典序++
            }
        }
        return ans;
    }
}
