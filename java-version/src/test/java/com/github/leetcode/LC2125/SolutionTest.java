package com.github.leetcode.LC2125;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numberOfBeams() {
        Solution solution = new Solution();
        assertEquals(8, solution.numberOfBeams(new String[]{"011001","000000","010100","001000"}));
        assertEquals(0, solution.numberOfBeams(new String[]{"000","111","000"}));
    }
}
