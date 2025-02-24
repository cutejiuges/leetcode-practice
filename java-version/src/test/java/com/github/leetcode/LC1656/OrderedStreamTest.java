package com.github.leetcode.LC1656;

import com.github.leetcode.LC59.Solution;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class OrderedStreamTest {

    @Test
    void insert() {
        OrderedStream stream = new OrderedStream(5);
        System.out.println(stream.insert(3, "ccccc"));
        System.out.println(stream.insert(1, "aaaaa"));
        System.out.println(stream.insert(2, "bbbbb"));
        System.out.println(stream.insert(5, "eeeee"));
        System.out.println(stream.insert(4, "ddddd"));
    }
}
