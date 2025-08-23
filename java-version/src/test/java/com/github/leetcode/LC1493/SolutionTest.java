package com.github.leetcode.LC1493;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class SolutionTest {

    @Test
    void testLongestSubarray() {
        Solution solution = new Solution();
        assertEquals(3, solution.longestSubarray(new int[]{1,1,0,1}));
        assertEquals(5, solution.longestSubarray(new int[]{0,1,1,1,0,1,1,0,1}));
        assertEquals(2, solution.longestSubarray(new int[]{1,1,1}));
    }
}
