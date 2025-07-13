package com.github.leetcode.LC2410;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void matchPlayersAndTrainers() {
        Solution solution = new Solution();
        assertEquals(2, solution.matchPlayersAndTrainers(new int[]{4,7,9}, new int[]{8,2,5,8}));
        assertEquals(1, solution.matchPlayersAndTrainers(new int[]{1,1,1}, new int[]{10}));
    }
}
