package com.github.leetcode.LC3020;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximumLength() {
        Solution solution = new Solution();
        assertEquals(3, solution.maximumLength(new int[]{5, 4, 1, 2, 2}));
        assertEquals(1, solution.maximumLength(new int[]{1, 3, 2, 4}));
    }
}
