package com.github.leetcode.LC3637;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isTrionic() {
        Solution solution = new Solution();
        assertTrue(solution.isTrionic(new int[]{1,3,5,4,2,6}));
        assertFalse(solution.isTrionic(new int[]{2,1,3}));
        assertFalse(solution.isTrionic(new int[]{5,6,4,6,8,8,7}));
    }
}
