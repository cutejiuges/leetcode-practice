package com.github.leetcode.LC2874;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximumTripletValue() {
        Solution solution = new Solution();
        assertEquals(77, solution.maximumTripletValue(new int[]{12,6,1,2,7}));
        assertEquals(133, solution.maximumTripletValue(new int[]{1,10,3,4,19}));
        assertEquals(0, solution.maximumTripletValue(new int[]{1,2,3}));
    }
}
