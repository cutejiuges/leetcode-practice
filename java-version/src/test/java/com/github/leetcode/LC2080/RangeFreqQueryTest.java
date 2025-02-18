package com.github.leetcode.LC2080;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class RangeFreqQueryTest {
    @Test
    public void TestQuery() {
        RangeFreqQuery query = new RangeFreqQuery(new int[]{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56});
        assertEquals(1, query.query(1, 2, 4));
        assertEquals(2, query.query(0, 11, 33));
    }
}
