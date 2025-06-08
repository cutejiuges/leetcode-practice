class Solution:
    @staticmethod
    def lexicalOrder(n: int) -> list[int]:
        ans = []
        num = 1
        for i in range(n):
            ans.append(num)
            if num * 10 <= n:
                num *= 10
            else:
                while num % 10 == 9 or num >= n:
                    num //= 10
                num += 1
        return ans
