package com.github.leetcode.LC440;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findKthNumber() {
        Solution solution = new Solution();
        assertEquals(10, solution.findKthNumber(13, 2));
        assertEquals(1, solution.findKthNumber(1, 1));
        assertEquals(416126219, solution.findKthNumber(681692778, 351251360));
    }
}
