import math


class Solution:
    @staticmethod
    def countTriples(n: int) -> int:
        ans = 0
        for a in range(1, n):
            for b in range(1, a):
                c = a * a + b * b
                if a * a + b * b > n * n:
                    break
                rt = int(math.sqrt(c))
                if rt ** 2 == c:
                    ans += 1
        return ans * 2
