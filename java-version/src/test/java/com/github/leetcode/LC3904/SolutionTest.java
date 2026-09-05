package com.github.leetcode.LC3904;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void firstStableIndex() {
        Solution solution = new Solution();
        assertEquals(3, solution.firstStableIndex(new int[]{5, 0, 1, 4}, 3));
        assertEquals(-1, solution.firstStableIndex(new int[]{3, 2, 1}, 1));
        assertEquals(0, solution.firstStableIndex(new int[]{0}, 0));
    }
}
