package com.github.leetcode.LC1323;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximum69Number() {
        Solution solution = new Solution();
        assertEquals(9969, solution.maximum69Number(9669));
        assertEquals(9999, solution.maximum69Number(9996));
        assertEquals(9999, solution.maximum69Number(9999));
    }
}
