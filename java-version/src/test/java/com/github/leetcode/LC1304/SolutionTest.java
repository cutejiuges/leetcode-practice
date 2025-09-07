package com.github.leetcode.LC1304;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void sumZero() {
        Solution solution = new Solution();
        int[] ans = solution.sumZero(5);
        int sum = Arrays.stream(ans).sum();
        assertEquals(0, sum);
        ans = solution.sumZero(3);
        sum = Arrays.stream(ans).sum();
        assertEquals(0, sum);
        ans = solution.sumZero(1);
        sum = Arrays.stream(ans).sum();
        assertEquals(0, sum);
        ans = solution.sumZero(6);
        sum = Arrays.stream(ans).sum();
        assertEquals(0, sum);
    }
}
