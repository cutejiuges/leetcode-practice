package com.github.leetcode.LC790;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numTilings() {
        Solution solution = new Solution();
        assertEquals(5, solution.numTilings(3));
        assertEquals(1, solution.numTilings(1));
    }
}
