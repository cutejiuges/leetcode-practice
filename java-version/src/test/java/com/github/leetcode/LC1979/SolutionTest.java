package com.github.leetcode.LC1979;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findGCD() {
        Solution solution = new Solution();
        assertEquals(2, solution.findGCD(new int[]{2,5,6,9,10}));
        assertEquals(1, solution.findGCD(new int[]{7,5,6,8,3}));
        assertEquals(3, solution.findGCD(new int[]{3,3}));
    }
}
