class Solution:
    def findEvenNumbers(self, digits: list[int]) -> list[int]:
        cnt = [0] * 10
        for _, digit in enumerate(digits):
            cnt[digit] += 1
        ans = []
        for i in range(100, 1000, 2):
            if self.__judge(i, cnt):
                ans.append(i)
        return ans

    @staticmethod
    def __judge(num: int, cnt: list[int]) -> bool:
        c = [0] * 10
        x = num
        while x > 0:
            cur = x % 10
            c[cur] += 1
            if c[cur] > cnt[cur]:
                return False
            x //= 10
        return True
