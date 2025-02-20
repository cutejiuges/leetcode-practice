package com.github.leetcode.LC2595;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestEvenOddBit() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{1, 2}, solution.evenOddBit(50));
        assertArrayEquals(new int[]{0, 1}, solution.evenOddBit(2));
    }
}
