package com.github.leetcode.LC3024;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void triangleType() {
        Solution solution = new Solution();
        assertEquals("equilateral", solution.triangleType(new int[]{3,3,3}));
        assertEquals("scalene", solution.triangleType(new int[]{3,4,5}));
    }
}
