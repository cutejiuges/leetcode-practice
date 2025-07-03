package com.github.leetcode.LC3304;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void kthCharacter() {
        Solution solution = new Solution();
        assertEquals('b', solution.kthCharacter(5));
        assertEquals('c', solution.kthCharacter(10));
    }
}
