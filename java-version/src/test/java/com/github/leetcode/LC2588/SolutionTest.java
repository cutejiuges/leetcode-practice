package com.github.leetcode.LC2588;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void beautifulSubarrays() {
        Solution solution = new Solution();
        assertEquals(2, solution.beautifulSubarrays(new int[]{4,3,1,2,4}));
        assertEquals(0, solution.beautifulSubarrays(new int[]{1,10,4}));
    }
}
