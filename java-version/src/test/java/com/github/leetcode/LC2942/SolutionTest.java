package com.github.leetcode.LC2942;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findWordsContaining() {
        Solution solution = new Solution();
        assertArrayEquals(new Object[]{0,1}, solution.findWordsContaining(new String[]{"leet","code"}, 'e').toArray());
        assertArrayEquals(new Object[]{0,2}, solution.findWordsContaining(new String[]{"abc","bcd","aaaa","cbc"}, 'a').toArray());
        assertArrayEquals(new Object[]{}, solution.findWordsContaining(new String[]{"abc","bcd","aaaa","cbc"}, 'z').toArray());
    }
}
