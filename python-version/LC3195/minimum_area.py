from typing import List


class Solution:
    @staticmethod
    def minimumArea(grid: List[List[int]]) -> int:
        if not grid or not grid[0]:
            return 0
        m, n = len(grid), len(grid[0])
        min_row, min_col = m, n
        max_row, max_col = 0, 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    min_row = min(min_row, i)
                    min_col = min(min_col, j)
                    max_row = max(max_row, i)
                    max_col = max(max_col, j)
        return (max_row - min_row + 1) * (max_col - min_col + 1)
