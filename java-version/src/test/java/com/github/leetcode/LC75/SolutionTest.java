package com.github.leetcode.LC75;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void sortColors() {
        Solution solution = new Solution();
        int[] ans = new int[]{2,0,2,1,1,0};
        solution.sortColors(ans);
        assertArrayEquals(new int[]{0,0,1,1,2,2}, ans);
        ans = new int[]{2,0,1};
        solution.sortColors(ans);
        assertArrayEquals(new int[]{0,1,2}, ans);
    }
}
