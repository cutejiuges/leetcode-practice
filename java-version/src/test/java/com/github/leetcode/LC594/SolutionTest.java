package com.github.leetcode.LC594;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findLHS() {
        Solution solution = new Solution();
        assertEquals(5, solution.findLHS(new int[]{1,3,2,2,5,2,3,7}));
        assertEquals(2, solution.findLHS(new int[]{1,2,3,4}));
        assertEquals(0, solution.findLHS(new int[]{1,1,1,1}));
    }
}
