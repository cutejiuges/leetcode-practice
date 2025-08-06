package com.github.leetcode.LC3479;

public class Solution {
    public int numOfUnplacedFruits(int[] fruits, int[] baskets) {
        SegmentTree segmentTree = new SegmentTree(baskets);
        int n = baskets.length, ans = 0;
        for (int fruit : fruits) {
            if (segmentTree.findFirstAndUpdate(1, 0, n - 1, fruit) < 0) {
                ans++;
            }
        }
        return ans;
    }
}

// 定义区间树结构
class SegmentTree {
    private final int[] max;

    public SegmentTree(int[] arr) {
        int n = arr.length;
        max = new int[2 << (32 - Integer.numberOfLeadingZeros(n))];
        build(arr, 1, 0, n - 1);
    }

    // 找到区间内第一个 >= x的数，更新为-1,返回这个数字的下标，如果没有的话返回-1
    public int findFirstAndUpdate(int o, int left, int right, int x) {
        if (max[o] < x) { // 区间内没有>=x的数字
            return -1;
        }
        if (left == right) {
            max[o] = -1; // 更新为-1, 标记当前位置已经被占用，不能放水果
            return left;
        }
        int mid = left + (right - left) / 2;
        int idx = findFirstAndUpdate(o * 2, left, mid, x); // 先递归左边子树
        if (idx == -1) { // 没找到再递归右子树
            idx = findFirstAndUpdate(o * 2 + 1, mid + 1, right, x);
        }
        maintain(o);
        return idx;
    }

    private void maintain(int o) {
        max[o] = Math.max(max[o * 2], max[o * 2 + 1]);
    }

    // 初始化区间树
    private void build(int[] arr, int o, int left, int right) {
        if (left == right) {
            max[o] = arr[left];
            return;
        }
        int mid = left + (right - left) / 2;
        build(arr, o * 2, left, mid);
        build(arr, o * 2 + 1, mid + 1, right);
        maintain(o);
    }
}
