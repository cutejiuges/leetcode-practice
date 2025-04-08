package com.github.leetcode.LC3375;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minOperations() {
        Solution solution = new Solution();
        assertEquals(2, solution.minOperations(new int[]{5,2,5,4,5}, 2));
        assertEquals(-1, solution.minOperations(new int[]{2,1,2}, 2));
        assertEquals(4, solution.minOperations(new int[]{9,7,5,3}, 1));
    }
}
