package com.github.leetcode.LC3633;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void earliestFinishTime() {
        Solution solution = new Solution();
        assertEquals(14, solution.earliestFinishTime(new int[]{5}, new int[]{3}, new int[]{1}, new int[]{10}));
        assertEquals(9, solution.earliestFinishTime(new int[]{2,8}, new int[]{4, 1}, new int[]{6}, new int[]{3}));
    }
}
