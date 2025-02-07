package com.github.leetcode.LC3270;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void testGenerateKey() {
        Solution solution = new Solution();
        assertEquals(0, solution.generateKey(1, 10, 1000));
        assertEquals(777, solution.generateKey(987, 879, 798));
        assertEquals(1, solution.generateKey(1, 2, 3));
    }
}
