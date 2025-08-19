from typing import List


class Solution:
    @staticmethod
    def zeroFilledSubarray(nums: List[int]) -> int:
        cnt, ans = 0, 0
        for num in nums:
            if num == 0:
                cnt += 1
                ans += cnt
            else:
                cnt = 0
        return ans
