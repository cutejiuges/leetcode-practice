package com.github.leetcode.LC2327;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void peopleAwareOfSecret() {
        Solution solution = new Solution();
        assertEquals(5, solution.peopleAwareOfSecret(6,2,4));
        assertEquals(6, solution.peopleAwareOfSecret(4,1,3));
    }
}
