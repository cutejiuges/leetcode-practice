package com.github.leetcode.LC3306;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countOfSubstrings() {
        Solution solution = new Solution();
        assertEquals(0, solution.countOfSubstrings("aeioqq", 1));
        assertEquals(1, solution.countOfSubstrings("aeiou", 0));
        assertEquals(3, solution.countOfSubstrings("ieaouqqieaouqq", 1));
    }
}
