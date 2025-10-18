package com.github.leetcode.LC3397;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxDistinctElements() {
        Solution solution = new Solution();
        assertEquals(6, solution.maxDistinctElements(new int[]{1,2,2,3,3,4}, 2));
        assertEquals(3, solution.maxDistinctElements(new int[]{4,4,4,4}, 1));
    }
}
