package com.github.leetcode.LC1706;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestFindBall() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{1,-1,-1,-1,-1}, 
            solution.findBall(new int[][]{{1,1,1,-1,-1},{1,1,1,-1,-1},{-1,-1,-1,1,1},{1,1,1,1,-1},{-1,-1,-1,-1,-1}}));
        assertArrayEquals(new int[]{-1}, solution.findBall(new int[][]{{-1}}));
        assertArrayEquals(new int[]{0,1,2,3,4,-1}, 
            solution.findBall(new int[][]{{1,1,1,1,1,1},{-1,-1,-1,-1,-1,-1},{1,1,1,1,1,1},{-1,-1,-1,-1,-1,-1}}));
    }
}
