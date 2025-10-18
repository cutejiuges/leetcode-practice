import math
from typing import List


class Solution:
    @staticmethod
    def maxDistinctElements(nums: List[int], k: int) -> int:
        ans = 0
        pre = -math.inf
        for num in nums:
            v = min(max(pre + 1, num - k), num + k)
            if v > pre:
                pre = v
                ans += 1
        return ans
