class Solution:
    def __init__(self):
        self.__nums = None
        self.__k = None
        self.__map = None
        self.__ans = -1

    def beautifulSubsets(self, nums: list[int], k: int) -> int:
        self.__nums = nums
        self.__k = k
        self.__map = {}
        self.__ans = -1
        self.backtrack(0)
        return self.__ans

    def backtrack(self, idx: int):
        if idx == len(self.__nums):
            self.__ans += 1
            return
        # 不选择当前元素，直接往下走
        self.backtrack(idx+1)
        # 如果能选就选中
        cur = self.__nums[idx]
        if self.__map.get(cur - self.__k, 0) == 0 and self.__map.get(cur + self.__k, 0) == 0:
            self.__map[cur] = self.__map.get(cur, 0) + 1
            self.backtrack(idx+1)
            # 恢复现场
            self.__map[cur] = self.__map.get(cur, 0) - 1
