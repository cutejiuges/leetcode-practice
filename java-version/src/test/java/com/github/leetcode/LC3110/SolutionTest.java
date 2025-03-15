package com.github.leetcode.LC3110;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void scoreOfString() {
        Solution solution = new Solution();
        assertEquals(13, solution.scoreOfString("hello"));
        assertEquals(50, solution.scoreOfString("zaz"));
    }
}
