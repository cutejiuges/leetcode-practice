package com.github.leetcode.LC2502;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class AllocatorTest {
    @Test
    public void TestAllocator() {
        Allocator allocator = new Allocator(10);
        System.out.println(allocator.allocate(1, 1));
        System.out.println(allocator.allocate(1, 2));
        System.out.println(allocator.allocate(1, 3));
        System.out.println(allocator.freeMemory(2));
        System.out.println(allocator.allocate(3, 4));
        System.out.println(allocator.allocate(1, 1));
        System.out.println(allocator.allocate(1, 1));
        System.out.println(allocator.freeMemory(1));
        System.out.println(allocator.allocate(10, 2));
        System.out.println(allocator.freeMemory(7));
    }
}
