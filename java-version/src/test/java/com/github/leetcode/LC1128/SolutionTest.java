package com.github.leetcode.LC1128;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numEquivDominoPairs() {
        Solution solution = new Solution();
        assertEquals(1, solution.numEquivDominoPairs(new int[][]{{1,2}, {2,1}, {3,4}, {5,6}}));
        assertEquals(3, solution.numEquivDominoPairs(new int[][]{{1,2}, {1,2}, {1,1}, {1,2}, {2,2}}));
    }
}
