class Solution:
    def beautifulSubarrays(self, nums: list[int]) -> int:
        map = {0:1}
        mask, ans = 0, 0
        for m in nums:
            mask ^= m
            ans += map.get(mask, 0)
            map[mask] = map.get(mask, 0) + 1
        return ans