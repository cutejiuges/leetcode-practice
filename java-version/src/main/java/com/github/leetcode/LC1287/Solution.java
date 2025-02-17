package com.github.leetcode.LC1287;

public class Solution {
    public int findSpecialInteger(int[] arr) {
        int num = arr[0], cnt = 0;
        for(int i = 0; i < arr.length; i++) {
            if(arr[i] == num) {
                cnt++;
                if(cnt * 4 > arr.length) {
                    return num;
                }
            } else {
                num = arr[i];
                cnt = 1;
            }
        }
        return -1;
    }
}
