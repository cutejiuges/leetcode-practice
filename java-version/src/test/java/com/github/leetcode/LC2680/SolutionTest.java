package com.github.leetcode.LC2680;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximumOr() {
        Solution solution = new Solution();
        assertEquals(30, solution.maximumOr(new int[]{12, 9}, 1));
        assertEquals(35, solution.maximumOr(new int[]{8, 1, 2}, 2));
    }
}
