MOD = 10 ** 9 + 7
pow2 = [1] * 100000
for i in range(1, 100000):
    pow2[i] = pow2[i - 1] * 2 % MOD


class Solution:
    @staticmethod
    def numSubseq(nums: list[int], target: int) -> int:
        nums.sort()
        ans = 0
        low, high = 0, len(nums) - 1
        while low <= high:
            if nums[low] + nums[high] <= target:
                ans += pow2[high - low] % MOD
                low += 1
            else:
                high -= 1
        return ans % MOD
