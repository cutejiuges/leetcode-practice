class Solution:
    def checkPowersOfThree(self, n: int) -> bool:
        return self.__powerSumOfX(n, 3)

    @staticmethod
    def __powerSumOfX(num: int, n: int) -> bool:
        if num < 1:
            return False
        while num != 0:
            if num % n > 1:
                return False
            num //= n
        return True
