package com.github.leetcode.LC165;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void compareVersion() {
        Solution solution = new Solution();
        assertEquals(-1, solution.compareVersion("1.2", "1.10"));
        assertEquals(0, solution.compareVersion("1.01", "1.001"));
        assertEquals(0, solution.compareVersion("1.0", "1.0.0.0"));
    }
}
