package com.github.leetcode.LC2537;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countGood() {
        Solution solution = new Solution();
        assertEquals(1, solution.countGood(new int[]{1,1,1,1,1}, 10));
        assertEquals(4, solution.countGood(new int[]{3,1,4,3,2,2,4}, 2));
    }
}
