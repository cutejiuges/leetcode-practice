package com.github.leetcode.LC3201;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximumLength() {
        Solution solution = new Solution();
        assertEquals(4, solution.maximumLength(new int[]{1,2,3,4}));
        assertEquals(6, solution.maximumLength(new int[]{1,2,1,1,2,1,2}));
        assertEquals(2, solution.maximumLength(new int[]{1,3}));
    }
}
