package com.github.leetcode.LC2492;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    int ans = Integer.MAX_VALUE;

    public int minScore(int n, int[][] roads) {
        boolean[] visited = new boolean[n + 1];
        List<Edge>[] graph = new ArrayList[n + 1];
        for (int i = 0; i < n + 1; i++) {
            graph[i] = new ArrayList<>();
        }
        for (int[] road : roads) {
            int u = road[0], v = road[1], dist = road[2];
            graph[u].add(new Edge(v, dist));
            graph[v].add(new Edge(u, dist));
        }

        dfs(1, graph, visited);
        return ans;
    }

    private void dfs(int u, List<Edge>[] graph, boolean[] vis) {
        if (!vis[u]) {
            vis[u] = true;
        }
        for (Edge e : graph[u]) {
            if (e.dist < ans) {
                ans = e.dist;
            }
            if (!vis[e.v]) {
                dfs(e.v, graph, vis);
            }
        }
    }
}

class Edge {
    int v; // 可达节点编号
    int dist; // 到可达节点的距离
    public Edge(int v, int dist) {
        this.v = v;
        this.dist = dist;
    }
}
