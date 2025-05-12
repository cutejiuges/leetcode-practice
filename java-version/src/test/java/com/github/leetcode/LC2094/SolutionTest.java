package com.github.leetcode.LC2094;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findEvenNumbers() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{102,120,130,132,210,230,302,310,312,320}, solution.findEvenNumbers(new int[]{2,1,3,0}));
        assertArrayEquals(new int[]{222,228,282,288,822,828,882}, solution.findEvenNumbers(new int[]{2,2,8,8,2}));
        assertArrayEquals(new int[]{}, solution.findEvenNumbers(new int[]{3,7,5}));
    }
}
