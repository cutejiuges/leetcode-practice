class Solution:
    def __init__(self):
        self.__n = 0
        self.__ans = 0

    def __dfs(self, val: int, idx: int, nums: list[int]):
        if idx == self.__n:  # 递归终止条件
            self.__ans += val
            return
        # 选择当前数字
        self.__dfs(val ^ nums[idx], idx + 1, nums)
        # 跳过当前数字
        self.__dfs(val, idx + 1, nums)

    def subsetXORSum(self, nums: list[int]) -> int:
        self.__n = len(nums)
        self.__ans = 0
        self.__dfs(0, 0, nums)
        return self.__ans
