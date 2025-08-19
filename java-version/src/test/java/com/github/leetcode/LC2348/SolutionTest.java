package com.github.leetcode.LC2348;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void zeroFilledSubarray() {
        Solution solution = new Solution();
        assertEquals(6, solution.zeroFilledSubarray(new int[]{1,3,0,0,2,0,0,4}));
        assertEquals(9, solution.zeroFilledSubarray(new int[]{0,0,0,2,0,0}));
        assertEquals(0, solution.zeroFilledSubarray(new int[]{2,10,2019}));
    }
}
