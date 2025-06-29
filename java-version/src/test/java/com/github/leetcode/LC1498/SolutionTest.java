package com.github.leetcode.LC1498;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numSubseq() {
        Solution solution = new Solution();
        assertEquals(4, solution.numSubseq(new int[]{3,5,6,7}, 9));
        assertEquals(6, solution.numSubseq(new int[]{3,3,6,8}, 10));
        assertEquals(61, solution.numSubseq(new int[]{2,3,3,4,6,7}, 12));
    }
}
