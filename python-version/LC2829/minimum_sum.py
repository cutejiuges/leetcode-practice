class Solution:
    def minimumSum(self, n: int, k: int) -> int:
        left_sum = self.__sum_arithmetic_sequence(1, 1, min(n, (k >> 1)))
        right_sum = self.__sum_arithmetic_sequence(k, 1, n - (k >> 1))
        if n <= (k >> 1):
            return left_sum
        return left_sum + right_sum

    @staticmethod
    def __sum_arithmetic_sequence(a1: int, d: int, n: int) -> int:
        return n * a1 + ((n - 1) * n * d >> 1)
