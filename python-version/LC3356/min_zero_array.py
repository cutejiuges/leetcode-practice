class Solution:
    def minZeroArray(self, nums: list[int], queries: list[list[int]]) -> int:
        q = len(queries)
        low, high = 0, q+1
        while low < high:
            mid = low + (high - low) // 2
            if self.__check(nums, queries, mid):
                high = mid
            else:
                low = mid + 1
        return low if low <= q else -1

    @staticmethod
    def __check(nums: list[int], queries: list[list[int]], k: int) -> bool:
        n = len(nums)
        diff_array = [0] * n
        for query in queries[:k]:
            diff_array[query[0]] += query[2]
            if query[1] < n - 1:
                diff_array[query[1] + 1] -= query[2]

        sum_diff = 0
        for i, diff in enumerate(diff_array):
            sum_diff += diff
            if nums[i] > sum_diff:
                return False
        return True
        