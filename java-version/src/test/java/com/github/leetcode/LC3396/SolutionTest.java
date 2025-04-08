package com.github.leetcode.LC3396;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumOperations() {
        Solution solution = new Solution();
        assertEquals(2, solution.minimumOperations(new int[]{1,2,3,4,2,3,3,5,7}));
        assertEquals(2, solution.minimumOperations(new int[]{4,5,6,4,4}));
        assertEquals(0, solution.minimumOperations(new int[]{6,7,8,9}));
    }
}
