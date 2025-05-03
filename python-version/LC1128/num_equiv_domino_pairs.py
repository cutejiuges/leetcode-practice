class Solution:
    @staticmethod
    def numEquivDominoPairs(dominoes: list[list[int]]) -> int:
        visited = [[0] * 10 for _ in range(10)]
        ans = 0
        for domino in dominoes:
            x, y = min(domino), max(domino)
            ans += visited[x][y]
            visited[x][y] += 1
        return ans
