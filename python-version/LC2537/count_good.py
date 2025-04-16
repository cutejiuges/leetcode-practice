class Solution:
    def countGood(self, nums: list[int], k: int) -> int:
        ans = 0
        left, pairs = 0, 0
        cnt = {}
        for num in nums:
            pairs += cnt.get(num, 0)
            cnt[num] = cnt.get(num, 0) + 1
            while pairs >= k:
                pairs -= cnt[nums[left]] - 1
                cnt[nums[left]] = cnt[nums[left]] - 1
                left += 1
            ans += left
        return ans
