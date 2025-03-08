package com.github.leetcode.LC2234;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximumBeauty() {
        Solution solution = new Solution();
        assertEquals(14, solution.maximumBeauty(new int[]{1,3,1,1}, 7, 6, 12, 1));
        assertEquals(30, solution.maximumBeauty(new int[]{2,4,5,3}, 10, 5, 2, 6));
    }
}
