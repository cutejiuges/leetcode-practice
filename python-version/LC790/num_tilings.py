class Solution:
    @staticmethod
    def numTilings(n: int) -> int:
        mod = 10 ** 9 + 7
        dp = [[0] * 4 for _ in range(n + 1)]
        dp[1][0] = dp[1][1] = 1
        for i in range(2, n + 1):
            dp[i][0] = dp[i-1][1]
            cur = 0
            for j in range(4):
                cur = (cur + dp[i-1][j]) % mod
            dp[i][1] = cur
            dp[i][2] = (dp[i-1][0] + dp[i-1][3]) % mod
            dp[i][3] = (dp[i-1][0] + dp[i-1][2]) % mod
        return dp[n][1]
