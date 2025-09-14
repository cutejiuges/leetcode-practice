package com.github.leetcode.LC1733;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumTeachings() {
        Solution solution = new Solution();
        assertEquals(1, solution.minimumTeachings(2, new int[][]{{1}, {2}, {1,2}}, new int[][]{{1,2}, {1,3}, {2,3}}));
        assertEquals(2, solution.minimumTeachings(3, new int[][]{{2}, {1,3}, {1,2}, {3}}, new int[][]{{1,4}, {1,2}, {3,4}, {2,3}}));
    }
}
