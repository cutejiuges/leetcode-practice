class Solution:
    @staticmethod
    def maxAdjacentDistance(nums: list[int]) -> int:
        n = len(nums)
        ans = abs(nums[n-1] - nums[0])
        for i in range(1, n):
            ans = max(ans, abs(nums[i] - nums[i-1]))
        return ans
