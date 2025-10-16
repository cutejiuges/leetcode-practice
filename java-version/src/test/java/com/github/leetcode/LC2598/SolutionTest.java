package com.github.leetcode.LC2598;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findSmallestInteger() {
        Solution solution = new Solution();
        assertEquals(4, solution.findSmallestInteger(new int[]{1,-10,7,13,6,8}, 5));
        assertEquals(2, solution.findSmallestInteger(new int[]{1,-10,7,13,6,8}, 7));
    }
}
