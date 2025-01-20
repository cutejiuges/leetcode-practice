import math


class Solution:
    def findClosestNumber(self, nums: list[int]) -> int:
        ans = math.inf
        for x in nums:
            if math.fabs(x) < math.fabs(ans):
                ans = x
            elif math.fabs(x) == math.fabs(ans):
                ans = max(ans, x)
        return ans
    