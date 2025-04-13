package com.github.leetcode.LC1922;

public class Solution {
    //统计每一个数位可能的数字放置情况
    //偶数位每个位置可以放0,2,4,6,8这5种情况，奇数位可以放置2,3,5,7这4个数字
    //只需要分析数字有多少个奇数位和偶数位，利用快速幂即可完成运算
    //长度为n的字符串，偶数位有n/2向上取整个偶数位，n/2向下取整个奇数位
    public int countGoodNumbers(long n) {
        int mod = (int) 1e9+7;
        return (int) (quickPower(5, (n+1) / 2, mod) * quickPower(4, n / 2, mod) % mod);
    }

    private long quickPower(long a, long b, int p) {
        long ans = 1;
        while (b > 0) {
            if ((b & 1) == 1) {
                ans = ans * a % p;
            }
            a = a * a % p;
            b >>>= 1;
        }
        return ans;
    }
}
