package com.github.LC1561;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestMaxCoins() {
        Solution solution = new Solution();
        assertEquals(9, solution.maxCoins(new int[]{2,4,1,2,7,8}));
        assertEquals(4, solution.maxCoins(new int[]{2,4,5}));
        assertEquals(18, solution.maxCoins(new int[]{9,8,7,6,5,1,2,3,4}));
    }
}
