package com.github.leetcode.LC2829;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumSum() {
        Solution solution = new Solution();
        assertEquals(18, solution.minimumSum(5, 4));
        assertEquals(3, solution.minimumSum(2, 6));
    }
}
