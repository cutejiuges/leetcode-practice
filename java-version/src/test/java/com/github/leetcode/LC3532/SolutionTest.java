package com.github.leetcode.LC3532;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void pathExistenceQueries() {
        Solution solution = new Solution();
        assertArrayEquals(new boolean[]{true,false},
            solution.pathExistenceQueries(2, new int[]{1,3}, 1, new int[][]{{0,0},{0,1}}));
        assertArrayEquals(new boolean[]{false,false,true,true},
            solution.pathExistenceQueries(4, new int[]{2,5,6,8}, 2, new int[][]{{0,1},{0,2},{1,3},{2,3}}));
    }
}
