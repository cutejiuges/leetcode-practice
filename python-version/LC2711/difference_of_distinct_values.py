class Solution:
    def differenceOfDistinctValues(self, grid: list[list[int]]) -> list[list[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * n for _ in range(m)]
        for i, row in enumerate(grid):
            for j, val in enumerate(row):
                bottom_right_set = set()
                x, y = i+1, j+1
                while x < m and y < n:
                    bottom_right_set.add(grid[x][y])
                    x += 1
                    y += 1
                top_left_set = set()
                x, y = i-1, j-1
                while x >= 0 and y >= 0:
                    top_left_set.add(grid[x][y])
                    x -= 1
                    y -= 1
                ans[i][j] = abs(len(bottom_right_set) - len(top_left_set))
        return ans
