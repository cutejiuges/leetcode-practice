class Solution:
    def maxCoins(self, piles: list[int]) -> int:
        piles.sort()
        n, round = len(piles), len(piles) // 3
        ans, idx = 0, n - 2
        for _ in range(round):
            ans += piles[idx]
            idx -= 2
        return ans
