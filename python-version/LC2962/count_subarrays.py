class Solution:
    @staticmethod
    def countSubarrays(nums: list[int], k: int) -> int:
        mx = max(nums)
        left, ans, cnt_max = 0, 0, 0
        for num in nums:
            if num == mx:
                cnt_max += 1
            while cnt_max >= k:
                if nums[left] == mx:
                    cnt_max -= 1
                left += 1
            ans += left
        return ans
