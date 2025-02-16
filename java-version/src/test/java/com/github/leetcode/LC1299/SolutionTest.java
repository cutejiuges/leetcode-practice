package com.github.leetcode.LC1299;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestReplaceElements() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{18,6,6,6,1,-1}, solution.replaceElements(new int[]{17,18,5,4,6,1}));
        assertArrayEquals(new int[]{-1}, solution.replaceElements(new int[]{400}));
    }
}
