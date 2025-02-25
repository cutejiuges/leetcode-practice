package com.github.leetcode.LC2502;

public class Allocator {
    private int n;
    private int[] memory;

    public Allocator(int n) {
        this.n = n;
        this.memory = new int[n];
    }

    public int allocate(int size, int mID) {
        int cnt = 0;
        for (int i = 0; i < this.n; i++) {
            if (this.memory[i] != 0) {
                cnt = 0;
            } else {
                cnt++;
                if (cnt == size) {
                    for (int j = i - cnt + 1; j <= i; j++) {
                        this.memory[j] = mID;
                    }
                    return i-cnt+1;
                }
            }
        }
        return -1;
    }

    public int freeMemory(int mID) {
        int cnt = 0;
        for (int i = 0; i < this.n; i++) {
            if (this.memory[i] == mID) {
                this.memory[i] = 0;
                cnt++;
            }
        }
        return cnt;
    }
}
