from typing import List


class Solution:
    @staticmethod
    def hasIncreasingSubarrays(nums: List[int], k: int) -> bool:
        n, cur, pre, i = len(nums), 0, 0, 0
        while i < n:
            j = i
            while i + 1 < n and nums[i] < nums[i+1]:
                i += 1
            cur = i - j + 1
            if cur >= 2*k or (pre >= k and cur >= k):
                return True
            pre = cur
            i += 1
        return False
