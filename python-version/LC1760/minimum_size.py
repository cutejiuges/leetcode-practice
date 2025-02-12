class Solution:
    def minimumSize(self, nums: list[int], maxOperations: int) -> int:
        mx = max(nums)
        low, high = 0, mx
        while low+1 < high:
            mid = (low+high)//2
            if self.check(nums, maxOperations, mid):
                high = mid
            else:
                low = mid
        return high
    
    def check(self, nums: list[int], maxOperations: int, mid: int) -> bool:
        cnt = 0
        for x in nums:
            cnt += (x-1)//mid
        return cnt <= maxOperations