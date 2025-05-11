package com.github.leetcode.LC1550;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void threeConsecutiveOdds() {
        Solution solution = new Solution();
        assertFalse(solution.threeConsecutiveOdds(new int[]{2,6,4,1}));
        assertTrue(solution.threeConsecutiveOdds(new int[]{1,2,34,3,4,5,7,23,12}));
    }
}
