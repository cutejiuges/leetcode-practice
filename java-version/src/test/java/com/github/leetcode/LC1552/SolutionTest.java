package com.github.leetcode.LC1552;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestMaxDistance() {
        Solution solution = new Solution();
        assertEquals(3, solution.maxDistance(new int[]{1,2,3,4,7}, 3));
        assertEquals(999999999, solution.maxDistance(new int[]{5,4,3,2,1,1000000000}, 2));
    }
}
