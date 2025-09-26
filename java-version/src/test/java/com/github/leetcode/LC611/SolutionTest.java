package com.github.leetcode.LC611;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void triangleNumber() {
        Solution solution = new Solution();
        assertEquals(3, solution.triangleNumber(new int[]{2,2,3,4}));
        assertEquals(4, solution.triangleNumber(new int[]{4,2,3,4}));
    }
}
