package com.github.leetcode.LC2109;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void addSpaces() {
        Solution solution = new Solution();
        assertEquals("Leetcode Helps Me Learn", solution.addSpaces("LeetcodeHelpsMeLearn", new int[]{8,13,15}));
        assertEquals("i code in py thon", solution.addSpaces("icodeinpython", new int[]{1,5,7,9}));
        assertEquals(" s p a c i n g", solution.addSpaces("spacing", new int[]{0,1,2,3,4,5,6}));
    }
}
