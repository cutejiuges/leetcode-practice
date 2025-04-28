class Solution:
    @staticmethod
    def countSubarrays(nums: list[int], k: int) -> int:
        ans = 0
        left = 0
        score = 0
        for i, num in enumerate(nums):
            score += num
            while score * (i - left + 1) >= k:
                score -= nums[left]
                left += 1
            ans += (i - left + 1)
        return ans
