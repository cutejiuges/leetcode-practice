package com.github.leetcode.LC1358;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numberOfSubstrings() {
        Solution solution = new Solution();
        assertEquals(10, solution.numberOfSubstrings("abcabc"));
        assertEquals(3, solution.numberOfSubstrings("aaacb"));
        assertEquals(1, solution.numberOfSubstrings("abc"));
    }
}
