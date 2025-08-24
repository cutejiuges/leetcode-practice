package com.github.leetcode.LC498;

/**
 * lc498题，对角线转换遍历，注意边界问题
 *
 * @author cutejiuge
 * @since 2025/8/25 上午6:53
 */
public class Solution {
    public int[] findDiagonalOrder(int[][] mat) {
        int row = mat.length, col = mat[0].length;
        int[] res = new int[row * col];
        int m = 0, n = 0, cnt = row + col - 1;
        int idx = 0;
        for (int i = 0; i < cnt; i++) {
            if ((i & 1) == 0) { // 偶数个对角线向右上方
                while (m >= 0 && n < col) { // 没越界的情况下正常移动即可
                    res[idx++] = mat[m--][n++];
                }
                if (n < col) { // 行越界了，换到下一行
                    m++;
                } else { // 如果是列越界了，需要切换下两行，前一列
                    m += 2;
                    n--;
                }
            } else { // 奇数个对角线向左下方
                while (m < row && n >= 0) { // 没越界的情况下正常移动即可
                    res[idx++] = mat[m++][n--];
                }
                if (m < row) { // 列越界，换到下一列
                    n++;
                } else { // 行越界，切换到下两列，前一行
                    n += 2;
                    m--;
                }
            }
        }
        return res;
    }
}
