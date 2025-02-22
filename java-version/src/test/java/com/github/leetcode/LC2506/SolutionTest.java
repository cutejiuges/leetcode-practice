package com.github.leetcode.LC2506;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {
    @Test
    void TestSimilarPairs() {
        Solution solution = new Solution();
        assertEquals(2, solution.similarPairs(new String[]{"aba","aabb","abcd","bac","aabc"}));
        assertEquals(3, solution.similarPairs(new String[]{"aabb","ab","ba"}));
        assertEquals(0, solution.similarPairs(new String[]{"nba","cba","dba"}));
    }
}
