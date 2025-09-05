package com.github.leetcode.LC2749;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void makeTheIntegerZero() {
        Solution solution = new Solution();
        assertEquals(3, solution.makeTheIntegerZero(3, -2));
        assertEquals(-1, solution.makeTheIntegerZero(5, 7));
    }
}
