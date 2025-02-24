package com.github.leetcode.LC1656;

import java.util.ArrayList;
import java.util.List;

public class OrderedStream {
    private String[] stream;
    private int ptr;

    public OrderedStream(int n) {
        this.stream = new String[n+1];
        this.ptr = 1;
    }

    public List<String> insert(int idKey, String value) {
        this.stream[idKey] = value;
        List<String> list = new ArrayList<>();
        while (this.ptr < this.stream.length && this.stream[this.ptr] != null) {
            list.add(this.stream[this.ptr++]);
        }
        return list;
    }
}
