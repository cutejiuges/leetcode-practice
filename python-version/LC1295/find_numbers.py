class Solution:
    def findNumbers(self, nums: list[int]) -> int:
        ans = 0
        for num in nums:
            ans += 1 if self.__count_10_bits(num) & 1 == 0 else 0
        return ans

    @staticmethod
    def __count_10_bits(num: int) -> int:
        x, ans = num, 0
        while x > 0:
            x //= 10
            ans += 1
        return ans
