package com.github.leetcode.LC2894;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void differenceOfSums() {
        Solution solution = new Solution();
        assertEquals(19, solution.differenceOfSums(10, 3));
        assertEquals(15, solution.differenceOfSums(5, 6));
        assertEquals(-15, solution.differenceOfSums(5, 1));
    }
}
