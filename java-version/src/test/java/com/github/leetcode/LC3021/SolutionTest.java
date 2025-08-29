package com.github.leetcode.LC3021;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void flowerGame() {
        Solution solution = new Solution();
        assertEquals(3, solution.flowerGame(3, 2));
        assertEquals(0, solution.flowerGame(0, 0));
    }
}
