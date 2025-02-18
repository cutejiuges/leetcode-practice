package com.github.leetcode.LC2080;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class RangeFreqQuery {
    // 用于存储某个值，以及该值所在的位序
    private Map<Integer, List<Integer>> occurrence;

    public RangeFreqQuery(int[] arr) {
        this.occurrence = new HashMap<>();
        // 遍历数组构建map
        for (int i = 0; i < arr.length; i++) {
            occurrence.putIfAbsent(arr[i], new ArrayList<>());
            occurrence.get(arr[i]).add(i);
        }
    }

    public int query(int left, int right, int value) {
        //查找元素对应的下标位序
        List<Integer> position = this.occurrence.getOrDefault(value, new ArrayList<>());
        //二分查找左边界和右边界
        int low = lowBound(position, left), high = highBound(position, right);
        return high - low;
    }

    //找到第一个大于等于左边界的位序
    private int lowBound(List<Integer> position, int low) {
        int left = 0, right = position.size()-1;
        while (left <= right) {
            int mid = left + ((right - left) >>> 1);
            if (position.get(mid) < low) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return left;
    }

    //找到第一个大于等于右边界的位序
    private int highBound(List<Integer> position, int high) {
        int left = 0, right = position.size()-1;
        while(left <= right) {
            int mid = left + ((right - left) >>> 1);
            if(position.get(mid) <= high) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return left;
    }
}
