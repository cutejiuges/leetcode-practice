package com.github.leetcode.LC2278;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void percentageLetter() {
        Solution solution = new Solution();
        assertEquals(33, solution.percentageLetter("foobar", 'o'));
        assertEquals(0, solution.percentageLetter("jjjj", 'k'));
    }
}
