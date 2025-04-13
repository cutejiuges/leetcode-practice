class Solution:
    def countGoodNumbers(self, n: int) -> int:
        mod = int(1e9) + 7
        return self.__quick_power(5, (n+1)//2, mod) * self.__quick_power(4, n//2, mod) % mod

    @staticmethod
    def __quick_power(a: int, b: int, p: int) -> int:
        ans = 1
        while b > 0:
            if b & 1 == 1:
                ans = ans * a % p
            a = a * a % p
            b >>= 1
        return ans
