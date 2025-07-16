from typing import List


class Solution:
    @staticmethod
    def maximumLength(nums: List[int]) -> int:
        res = 0
        patterns = [[0, 0], [0, 1], [1, 0], [1, 1]]
        for pattern in patterns:
            cnt = 0
            for num in nums:
                if num & 1 == pattern[cnt & 1]:
                    cnt += 1
            res = max(res, cnt)
        return res
