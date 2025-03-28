package com.github.leetcode.LC2716;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimizedStringLength() {
        Solution solution = new Solution();
        assertEquals(3, solution.minimizedStringLength("aaabc"));
        assertEquals(3, solution.minimizedStringLength("cbbd"));
        assertEquals(2, solution.minimizedStringLength("aaaddd"));
    }
}
