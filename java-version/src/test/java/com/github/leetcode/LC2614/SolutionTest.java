package com.github.leetcode.LC2614;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void diagonalPrime() {
        Solution solution = new Solution();
        assertEquals(11, solution.diagonalPrime(new int[][]{{1,2,3}, {5,6,7}, {9,10,11}}));
        assertEquals(17, solution.diagonalPrime(new int[][]{{1,2,3}, {5,17,7}, {9,11,10}}));
    }
}
