package com.github.leetcode.LC2145;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numberOfArrays() {
        Solution solution = new Solution();
        assertEquals(2, solution.numberOfArrays(new int[]{1,-3,4}, 1, 6));
        assertEquals(4, solution.numberOfArrays(new int[]{3,-4,5,1,-2}, -4, 5));
        assertEquals(0, solution.numberOfArrays(new int[]{4,-7,2}, 3, 6));
    }
}
