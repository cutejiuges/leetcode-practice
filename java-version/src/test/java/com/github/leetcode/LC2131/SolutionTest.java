package com.github.leetcode.LC2131;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void longestPalindrome() {
        Solution solution = new Solution();
        assertEquals(6, solution.longestPalindrome(new String[]{"lc","cl","gg"}));
        assertEquals(8, solution.longestPalindrome(new String[]{"ab","ty","yt","lc","cl","ab"}));
        assertEquals(2, solution.longestPalindrome(new String[]{"cc","ll","xx"}));
    }
}
