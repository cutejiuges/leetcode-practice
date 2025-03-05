package com.github.leetcode.LC1328;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void breakPalindrome() {
        Solution solution = new Solution();
        assertEquals("aaccba", solution.breakPalindrome("abccba"));
        assertEquals("", solution.breakPalindrome("a"));
    }
}
