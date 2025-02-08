class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        cur_size = 2
        for i in range(2, len(nums)):
            if nums[i] != nums[cur_size-2]:
                nums[cur_size] = nums[i]
                cur_size += 1
        return min(cur_size, len(nums))