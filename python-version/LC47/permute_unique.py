class Solution:
    def __init__(self):
        self.visited = []

    def backtrack(self, nums: list[int], ans: list[list[int]], idx: int, perm: list[int]):
        if idx == len(nums):
            ans.append(perm[:])
            return
        for i in range (len(nums)):
            if self.visited[i] or (i > 0 and not self.visited[i-1] and nums[i] == nums[i-1]):
                continue
            perm.append(nums[i])
            self.visited[i] = True
            self.backtrack(nums, ans, idx+1, perm)
            self.visited[i] = False
            del(perm[-1])
    
    def permuteUnique(self, nums: list[int]) -> list[list[int]]:
        nums.sort()
        ans = []
        perm = []
        self.visited = [False] * len(nums)
        self.backtrack(nums, ans, 0, perm)
        return ans