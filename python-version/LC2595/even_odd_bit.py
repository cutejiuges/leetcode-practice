class Solution:
    def evenOddBit(self, n: int) -> list[int]:
        ans = [0] * 2
        pos = 0
        while n > 0:
            ans[pos] += n & 1
            n >>= 1
            pos ^= 1
        return ans