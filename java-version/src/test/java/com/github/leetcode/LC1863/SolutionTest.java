package com.github.leetcode.LC1863;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void subsetXORSum() {
        Solution solution = new Solution();
        assertEquals(6, solution.subsetXORSum(new int[]{1,3}));
        assertEquals(28, solution.subsetXORSum(new int[]{5,1,6}));
        assertEquals(480, solution.subsetXORSum(new int[]{3,4,5,6,7,8}));
    }
}
