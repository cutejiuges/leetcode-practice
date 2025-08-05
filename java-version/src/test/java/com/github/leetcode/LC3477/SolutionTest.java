package com.github.leetcode.LC3477;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numOfUnplacedFruits() {
        Solution solution = new Solution();
        assertEquals(1, solution.numOfUnplacedFruits(new int[]{4,2,5}, new int[]{3,5,4}));
        assertEquals(0, solution.numOfUnplacedFruits(new int[]{3,6,1}, new int[]{6,4,7}));
    }
}
