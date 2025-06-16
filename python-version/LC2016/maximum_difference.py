class Solution:
    @staticmethod
    def maximumDifference(nums: list[int]) -> int:
        ans, pre_min = -1, nums[0]
        for num in nums[1:]:
            if num <= pre_min:
                pre_min = num
            else:
                ans = max(ans, num - pre_min)
        return ans
