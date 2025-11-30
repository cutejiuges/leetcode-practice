from typing import List


class Solution:
    @staticmethod
    def minSubarray(nums: List[int], p: int) -> int:
        sm = sum(nums)
        remain = sm % p
        if remain == 0:
            return 0

        n, ans = len(nums), len(nums)
        s = 0
        last = dict()
        last[s] = -1
        for i, num in enumerate(nums):
            s = (s + num) % p
            last[s] = i
            j = last.get((s - remain + p) % p, -n)
            ans = min(ans, i - j)
        return ans if ans < n else -1
