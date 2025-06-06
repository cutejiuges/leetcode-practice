package com.github.leetcode.LC2434;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void robotWithString() {
        Solution solution = new Solution();
        assertEquals("azz", solution.robotWithString("zza"));
        assertEquals("abc", solution.robotWithString("bac"));
        assertEquals("addb", solution.robotWithString("bdda"));
    }
}
