package com.github.leetcode.LC1920;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void buildArray() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{0,1,2,4,5,3}, solution.buildArray(new int[]{0,2,1,5,3,4}));
        assertArrayEquals(new int[]{4,5,0,1,2,3}, solution.buildArray(new int[]{5,0,1,2,3,4}));
    }
}
