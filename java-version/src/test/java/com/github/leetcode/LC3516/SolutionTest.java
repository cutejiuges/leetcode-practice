package com.github.leetcode.LC3516;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findClosest() {
        Solution solution = new Solution();
        assertEquals(1, solution.findClosest(2,7,4));
        assertEquals(2, solution.findClosest(2,5,4));
        assertEquals(0, solution.findClosest(1,5,3));
    }
}
