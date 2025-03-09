package com.github.leetcode.LC2070;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximumBeauty() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{2,4,5,5,6,6}, solution.maximumBeauty(new int[][]{{1,2},{3,2},{2,4},{5,6},{3,5}}, new int[]{1,2,3,4,5,6}));
        assertArrayEquals(new int[]{4}, solution.maximumBeauty(new int[][]{{1,2},{1,2},{1,3},{1,4}}, new int[]{1}));
        assertArrayEquals(new int[]{0}, solution.maximumBeauty(new int[][]{{10,1000}}, new int[]{5}));
    }
}

