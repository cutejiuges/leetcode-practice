package com.github.leetcode.LC2785;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void sortVowels() {
        Solution solution = new Solution();
        assertEquals("lEOtcede", solution.sortVowels("lEetcOde"));
        assertEquals("lYmpH", solution.sortVowels("lYmpH"));
    }
}
