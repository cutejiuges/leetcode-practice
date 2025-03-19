package com.github.leetcode.LC2610;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findMatrix() {
        Solution solution = new Solution();
        assertEquals(new ArrayList<List<Integer>>(){{
            add(new ArrayList<Integer>(){{add(1); add(2); add(3); add(4);}});
            add(new ArrayList<Integer>(){{add(1); add(3);}});
            add(new ArrayList<Integer>(){{add(1);}});
        }}, solution.findMatrix(new int[]{1,3,4,1,2,3,1}));
        assertEquals(new ArrayList<List<Integer>>(){{
            add(new ArrayList<Integer>(){{add(1); add(2); add(3); add(4);}});
        }}, solution.findMatrix(new int[]{1,2,3,4}));
    }
}
