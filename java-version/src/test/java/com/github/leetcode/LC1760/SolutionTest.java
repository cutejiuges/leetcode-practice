package com.github.leetcode.LC1760;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestMinimumSize() {
        Solution solution = new Solution();
        assertEquals(3, solution.minimumSize(new int[]{9}, 2));
        assertEquals(2, solution.minimumSize(new int[]{2,4,8,2}, 4));
        assertEquals(7, solution.minimumSize(new int[]{7, 17}, 2));
    }
}
