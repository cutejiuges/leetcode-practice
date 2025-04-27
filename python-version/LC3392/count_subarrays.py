class Solution:
    @staticmethod
    def countSubarrays(nums: list[int]) -> int:
        ans = 0
        n = len(nums)
        for i in range(1, n-1):
            if nums[i] == (nums[i-1]+ nums[i+1]) * 2:
                ans += 1
        return ans
