from typing import List


class Solution:
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        bins = self.__constructBinArray(n)
        return self.__processQueries(bins, queries)

    @staticmethod
    def __constructBinArray(n: int) -> List[int]:
        bins = []
        base = 1
        while n > 0:
            if n & 1 == 1:
                bins.append(base)
            base <<= 1
            n >>= 1
        return bins

    @staticmethod
    def __processQueries(bins: List[int], queries: List[List[int]]) -> List[int]:
        m = len(bins)
        mod = int(1e9+7)
        results = [[0 * m] * m for _ in range(m)]
        for i in range(m):
            cur = 1
            for j in range(i, m):
                cur = cur * bins[j] % mod
                results[i][j] = cur
        ans = [0] * len(queries)
        for i, query in enumerate(queries):
            ans[i] = results[query[0]][query[1]]
        return ans
