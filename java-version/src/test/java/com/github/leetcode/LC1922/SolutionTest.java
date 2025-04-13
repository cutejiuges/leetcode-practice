package com.github.leetcode.LC1922;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countGoodNumbers() {
        Solution solution = new Solution();
        assertEquals(5, solution.countGoodNumbers(1));
        assertEquals(400, solution.countGoodNumbers(4));
        assertEquals(564908303, solution.countGoodNumbers(50));
    }
}
