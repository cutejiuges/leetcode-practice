package com.github.leetcode.LC838;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void pushDominoes() {
        Solution solution = new Solution();
        assertEquals("RR.L", solution.pushDominoes("RR.L"));
        assertEquals("LL.RR.LLRRLL..", solution.pushDominoes(".L.R...LR..L.."));
    }
}
