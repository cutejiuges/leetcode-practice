package com.github.leetcode.LC3330;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void possibleStringCount() {
        Solution solution = new Solution();
        assertEquals(5, solution.possibleStringCount("abbcccc"));
        assertEquals(1, solution.possibleStringCount("abcd"));
        assertEquals(4, solution.possibleStringCount("aaaa"));
    }
}
