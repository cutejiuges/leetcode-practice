class Solution:
    def minSum(self, nums1: list[int], nums2: list[int]) -> int:
        res1, res2 = self.calculate(nums1), self.calculate(nums2)
        if res1[0] > res2[0] and res2[1] == 0:
            return -1
        if res2[0] > res1[0] and res1[1] == 0:
            return -1
        return max(res1[0], res2[0])

    @staticmethod
    def calculate(nums: list[int]) -> list[int]:
        sm, cnt0 = 0, 0
        for num in nums:
            sm += num if num > 0 else 1
            cnt0 += 1 if num == 0 else 0
        return [sm, cnt0]
