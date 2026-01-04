package com.github.leetcode.LC1390;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void sumFourDivisors() {
        Solution solution = new Solution();
        assertEquals(32, solution.sumFourDivisors(new int[]{21,4,7}));
        assertEquals(64, solution.sumFourDivisors(new int[]{21,21}));
        assertEquals(0, solution.sumFourDivisors(new int[]{1,2,3,4,5}));
    }
}
