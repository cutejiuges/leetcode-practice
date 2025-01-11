package com.github.LC3270;

public class Solution {
    public int generateKey(int num1, int num2, int num3) {
        int ans = 0;
        for(int pow = 1; num1 != 0 && num2 != 0 && num3 != 0; pow *= 10) {
            ans += Math.min(num1 % 10, Math.min(num2 % 10, num3 % 10)) * pow;
            num1 /= 10;
            num2 /= 10;
            num3 /= 10;
        }
        return ans;
    }
}
