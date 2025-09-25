package com.github.leetcode.LC120;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumTotal() {
        Solution solution = new Solution();
        assertEquals(11, solution.minimumTotal(Arrays.asList(
            Collections.singletonList(2),
            Arrays.asList(3, 4),
            Arrays.asList(6,5,7),
            Arrays.asList(4,1,8,3)
        )));
        assertEquals(-10, solution.minimumTotal(Arrays.asList(
            Collections.singletonList(-10)
        )));
    }
}
