package com.github.leetcode.LC1472;

import java.util.ArrayList;
import java.util.List;

public class BrowserHistory {
    private final List<String> urls;
    private int idx;

    public BrowserHistory(String homepage) {
        this.urls = new ArrayList<>();
        this.urls.add(homepage);
        this.idx = 0;
    }

    public void visit(String url) {
        while (this.urls.size() > this.idx+1) {
            this.urls.remove(this.urls.size() - 1);
        }
        this.urls.add(url);
        this.idx++;
    }

    public String back(int steps) {
        this.idx = Math.max(0, this.idx - steps);
        return this.urls.get(this.idx);
    }

    public String forward(int steps) {
        this.idx = Math.min(this.idx + steps, this.urls.size()-1);
        return this.urls.get(this.idx);
    }
}
