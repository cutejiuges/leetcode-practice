from typing import List


class Solution:
    @staticmethod
    def findSmallestInteger(nums: List[int], value: int) -> int:
        mp = [0] * value
        for num in nums:
            v = (num % value + value) % value
            mp[v] += 1

        mex = 0
        while mp[mex % value] > 0:
            mp[mex % value] -= 1
            mex += 1
        return mex
