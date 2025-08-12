package com.github.leetcode.LC2787;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numberOfWays() {
        Solution solution = new Solution();
        assertEquals(1, solution.numberOfWays(10, 2));
        assertEquals(2, solution.numberOfWays(4, 1));
    }
}
