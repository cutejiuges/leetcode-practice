package com.github.leetcode.LC3349;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void hasIncreasingSubarrays() {
        Solution solution = new Solution();
        assertTrue(solution.hasIncreasingSubarrays(Arrays.asList(2,5,7,8,9,2,3,4,3,1), 3));
        assertFalse(solution.hasIncreasingSubarrays(Arrays.asList(1,2,3,4,4,4,4,5,6,7),5));
    }
}
