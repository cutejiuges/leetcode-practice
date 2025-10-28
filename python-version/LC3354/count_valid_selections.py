from typing import List


class Solution:
    @staticmethod
    def countValidSelections(nums: List[int]) -> int:
        sm = sum(nums)
        cur, ans = 0, 0
        for num in nums:
            cur += num
            if num == 0 and cur * 2 == sm:
                ans += 2
            if num == 0 and abs(cur * 2 - sm) == 1:
                ans += 1
        return ans
