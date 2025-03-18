class Solution:
    def diagonalPrime(self, nums: list[list[int]]) -> int:
        primes = self.__init_primes()
        n = len(nums)
        ans = 0
        for i, _ in enumerate(nums):
            if primes[nums[i][i]]:
                ans = max(ans, nums[i][i])
            if primes[nums[i][n-i-1]]:
                ans = max(ans, nums[i][n-i-1])
        return ans

    @staticmethod
    def __init_primes() -> list[bool]:
        n = int(4*1e6+1)
        primes = [True] * n
        primes[0] = primes[1] = False
        for i in range(2, int(n ** 0.5)):
            if primes[i]:
                for j in range(i*i, n, i):
                    primes[j] = False
        return primes
