package com.github.leetcode.LC1399;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countLargestGroup() {
        Solution solution = new Solution();
        assertEquals(4, solution.countLargestGroup(13));
        assertEquals(2, solution.countLargestGroup(2));
        assertEquals(6, solution.countLargestGroup(15));
        assertEquals(5, solution.countLargestGroup(24));
    }
}
