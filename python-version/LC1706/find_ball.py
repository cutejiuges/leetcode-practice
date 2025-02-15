class Solution:
    def findBall(self, grid: list[list[int]]) -> list[int]:
        n = len(grid[0])
        ans = [0] * n
        for i in range(n):
            curCol = i
            for row in grid:
                dir = row[curCol]
                nextCol = curCol + dir
                if nextCol < 0 or nextCol >= n or row[nextCol] != dir:
                    curCol = -1
                    break
                curCol = nextCol
            ans[i] = curCol
        return ans