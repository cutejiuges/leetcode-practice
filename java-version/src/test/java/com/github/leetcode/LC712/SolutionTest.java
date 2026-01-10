package com.github.leetcode.LC712;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumDeleteSum() {
        Solution solution = new Solution();
        assertEquals(231, solution.minimumDeleteSum("sea", "eat"));
        assertEquals(403, solution.minimumDeleteSum("delete", "leet"));
    }
}
