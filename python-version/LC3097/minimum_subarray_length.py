import math

class Solution:
    def minimumSubarrayLength(self, nums: list[int], k: int) -> int:
        ans, n = math.inf, len(nums)
        left = 0
        bits = [0 for _ in range(32)]

        for right in range(n):
            for i in range(32):
                bits[i] += nums[right] >> i & 1
            while left <= right and self.valueOf(bits) >= k:
                ans = min(ans, right - left + 1)
                for i in range(32):
                    bits[i] -= nums[left] >> i & 1
                left += 1
        
        return -1 if ans == math.inf else ans
    
    def valueOf(self, bits: list[int]) -> int:
        n = len(bits)
        ans = 0
        for i in range(n):
            if bits[i] > 0:
                ans |= 1 << i
        return ans