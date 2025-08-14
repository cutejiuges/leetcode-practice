package com.github.leetcode.LC1780;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void checkPowersOfThree() {
        Solution solution = new Solution();
        assertTrue(solution.checkPowersOfThree(12));
        assertTrue(solution.checkPowersOfThree(91));
        assertFalse(solution.checkPowersOfThree(21));
    }
}
