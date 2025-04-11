package com.github.leetcode.LC2843;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countSymmetricIntegers() {
        Solution solution = new Solution();
        assertEquals(9, solution.countSymmetricIntegers(1, 100));
        assertEquals(4, solution.countSymmetricIntegers(1200, 1230));
    }
}
