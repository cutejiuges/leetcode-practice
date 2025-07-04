package com.github.leetcode.LC3307;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void kthCharacter() {
        Solution solution = new Solution();
        assertEquals('a', solution.kthCharacter(5, new int[]{0,0,0}));
        assertEquals('b', solution.kthCharacter(10, new int[]{0,1,0,1}));
    }
}
