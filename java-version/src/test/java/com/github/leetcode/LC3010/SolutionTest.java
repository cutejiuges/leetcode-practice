package com.github.leetcode.LC3010;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumCost() {
        Solution solution = new Solution();
        assertEquals(6, solution.minimumCost(new int[]{1, 2, 3, 12}));
        assertEquals(12, solution.minimumCost(new int[]{5,4,3}));
        assertEquals(12, solution.minimumCost(new int[]{10,3,1,1}));
    }
}
