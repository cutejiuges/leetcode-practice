package com.github.leetcode.LC2685;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public int countCompleteComponents(int n, int[][] edges) {
        // 构建邻接表
        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            graph.add(new ArrayList<>());
        }
        for (int[] e : edges) {
            graph.get(e[0]).add(e[1]);
            graph.get(e[1]).add(e[0]);
        }
        // 构建访问标记
        boolean[] visited = new boolean[n];

        int ans = 0;
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                int[] cnt = new int[2];
                dfs(i, graph, visited, cnt);
                if (cnt[0] * (cnt[0] - 1) == cnt[1]) {
                    ans++;
                }
            }
        }
        return ans;
    }

    private void dfs(int u, List<List<Integer>> graph, boolean[] visited, int[] cnt) {
        if (visited[u]) return;
        visited[u] = true;
        cnt[0]++; //节点数量加一
        cnt[1] += graph.get(u).size(); // 边的数量加上该节点出发的边
        for (int v : graph.get(u)) {
            if (!visited[v]) {
                dfs(v, graph, visited, cnt);
            }
        }
    }
}
