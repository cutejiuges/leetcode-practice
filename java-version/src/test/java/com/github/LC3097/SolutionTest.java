package com.github.LC3097;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestMinimumSubarrayLength() {
        Solution solution = new Solution();
        assertEquals(1, solution.minimumSubarrayLength(new int[]{1,2,3}, 2));
        assertEquals(3, solution.minimumSubarrayLength(new int[]{2,1,8}, 10));
        assertEquals(1, solution.minimumSubarrayLength(new int[]{1,2}, 0));
    }
}
