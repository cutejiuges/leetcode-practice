package com.github.LC2218;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestMaxValueOfCoins() {
        Solution solution = new Solution();
        List<List<Integer>> list = new ArrayList<>();
        list.add(new ArrayList<>(Arrays.asList(1, 100, 3)));
        list.add(new ArrayList<>(Arrays.asList(7, 8, 9)));
        assertEquals(101, solution.maxValueOfCoins(list, 2));

        list.clear();
        list.add(new ArrayList<>(Arrays.asList(100)));
        list.add(new ArrayList<>(Arrays.asList(100)));
        list.add(new ArrayList<>(Arrays.asList(100)));
        list.add(new ArrayList<>(Arrays.asList(100)));
        list.add(new ArrayList<>(Arrays.asList(100)));
        list.add(new ArrayList<>(Arrays.asList(100)));
        list.add(new ArrayList<>(Arrays.asList(1,1,1,1,1,1,700)));
        assertEquals(706, solution.maxValueOfCoins(list, 7));
    }
}
