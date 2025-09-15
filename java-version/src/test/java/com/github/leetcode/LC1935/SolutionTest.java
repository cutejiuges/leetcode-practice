package com.github.leetcode.LC1935;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void canBeTypedWords() {
        Solution solution = new Solution();
        assertEquals(1, solution.canBeTypedWords("hello world", "ad"));
        assertEquals(1, solution.canBeTypedWords("leet code", "lt"));
        assertEquals(0, solution.canBeTypedWords("leet code", "e"));
    }
}
