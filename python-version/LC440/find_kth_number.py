class Solution:
    def findKthNumber(self, n: int, k: int) -> int:
        cur = 1
        k -= 1
        while k > 0:
            step = self.get_steps(cur, n)
            if k >= step:
                k -= step
                cur += 1
            else:
                k -= 1
                cur *= 10
        return cur

    @staticmethod
    def get_steps(cur, n):
        step = 0
        first, last = cur, cur
        while first <= n:
            step += (min(last, n) - first) + 1
            first *= 10
            last = last * 10 + 9
        return step
