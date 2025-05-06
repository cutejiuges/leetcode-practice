class Solution:
    @staticmethod
    def buildArray(nums: list[int]) -> list[int]:
        length = len(nums)
        ans = [0] * length
        for i, _ in enumerate(nums):
            ans[i] = nums[nums[i]]
        return ans
