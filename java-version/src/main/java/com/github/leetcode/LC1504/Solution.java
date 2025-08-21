package com.github.leetcode.LC1504;

/**
 * lc1504，考虑使用遍历枚举矩形的右下角
 *
 * @author cutejiuge
 * @since 2025/8/21 下午10:41
 */
public class Solution {
    public int numSubmat(int[][] mat) {
        if (mat == null || mat.length == 0 || mat[0].length == 0) {
            return 0;
        }
        int m = mat.length, n = mat[0].length;
        int[][] temp = new int[m][n];
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (j == 0) { // 先统计横向的，如果是第一列，取当前位置的值
                    temp[i][j] = mat[i][j];
                } else { // 不是第一列，需要累积这一行之前的数量
                    temp[i][j] = mat[i][j] == 0 ? 0 : temp[i][j-1] + 1;
                }
                int cur = temp[i][j]; // 当前横向的已经统计完了，把纵向扩展的加上去
                for (int k = i; k >= 0; k--) { // 反方向向上找回去
                    cur = Math.min(cur, temp[k][j]); // 能构成多少个取决于上面构成了多少个
                    if (cur == 0) {
                        break;
                    }
                    ans += cur;
                }
            }
        }
        return ans;
    }
}
