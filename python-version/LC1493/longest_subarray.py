from typing import List


class Solution:
    @staticmethod
    def longestSubarray(nums: List[int]) -> int:
        cnt0, ans, left = 0, 0, 0
        for i, num in enumerate(nums):
            cnt0 += 1 if num == 0 else 0
            while cnt0 > 1:
                cnt0 -= 1 if nums[left] == 0 else 0
                left += 1
            ans = max(ans, i - left)
        return ans
