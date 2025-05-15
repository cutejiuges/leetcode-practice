package com.github.leetcode.LC2900;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void getLongestSubsequence() {
        Solution solution = new Solution();
        List<String> ans = solution.getLongestSubsequence(new String[]{"e","a","b"}, new int[]{0,0,1});
        List<String> expected = new ArrayList<String>(){{add("a");add("b");}};
        assertTrue(
            expected.size() == ans.size() &&
                expected.containsAll(ans) && ans.containsAll(expected)
        );
        ans = solution.getLongestSubsequence(new String[]{"a","b","c","d"}, new int[]{1,0,1,1});
        expected = new ArrayList<String>(){{add("a");add("b");add("c");}};
        assertTrue(
            expected.size() == ans.size() &&
                expected.containsAll(ans) && ans.containsAll(expected)
        );
    }
}
