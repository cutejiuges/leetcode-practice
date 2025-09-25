from typing import List


class Solution:
    @staticmethod
    def minimumTotal(triangle: List[List[int]]) -> int:
        if not triangle or not triangle[0]:
            return 0
        n = len(triangle)
        dp = [[0] * n for _ in range(n)]
        dp[0][0] = triangle[0][0]
        for i in range(1, n):
            dp[i][0] = dp[i-1][0] + triangle[i][0]
            for j in range(1, i):
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
            dp[i][i] = dp[i-1][i-1] + triangle[i][i]
        ans = min(dp[n-1])
        return ans
