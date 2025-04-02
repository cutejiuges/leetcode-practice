class Solution:
    def maximumTripletValue(self, nums: list[int]) -> int:
        n = len(nums)
        suf_max = [0]*(n+1)
        for i in range(n-1, 1, -1):
            suf_max[i] = max(nums[i], suf_max[i+1])

        pre_max = nums[0]
        ans = 0
        for i in range(1, n):
            ans = max(ans, (pre_max - nums[i]) * suf_max[i+1])
            pre_max = max(nums[i], pre_max)
        return ans
