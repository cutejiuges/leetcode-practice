from typing import List


class Solution:
    @staticmethod
    def triangleNumber(nums: List[int]) -> int:
        n, ans = len(nums), 0
        nums.sort()
        for i in range(n):
            k = i
            for j in range(i+1, n):
                while k + 1 < n and nums[k+1] < nums[i] + nums[j]:
                    k += 1
                ans += max(k-j, 0)
        return ans
