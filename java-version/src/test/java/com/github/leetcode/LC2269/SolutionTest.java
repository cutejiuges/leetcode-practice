package com.github.leetcode.LC2269;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void divisorSubstrings() {
        Solution solution = new Solution();
        assertEquals(2, solution.divisorSubstrings(240, 2));
        assertEquals(2, solution.divisorSubstrings(430043, 2));
    }
}
