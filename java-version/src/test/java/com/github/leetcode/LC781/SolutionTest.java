package com.github.leetcode.LC781;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numRabbits() {
        Solution solution = new Solution();
        assertEquals(5, solution.numRabbits(new int[]{1,1,2}));
        assertEquals(11, solution.numRabbits(new int[]{10,10,10}));
    }
}
