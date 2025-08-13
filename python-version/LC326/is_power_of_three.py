class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n < 3:
            return n == 1
        return n % 3 == 0 and self.isPowerOfThree(n // 3)
