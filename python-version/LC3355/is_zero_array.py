class Solution:
    @staticmethod
    def isZeroArray(nums: list[int], queries: list[list[int]]) -> bool:
        n = len(nums)
        diff_arr = [0] * n
        for query in queries:
            diff_arr[query[0]] += 1
            if query[1] < n - 1:
                diff_arr[query[1] + 1] -= 1

        sum_diff = 0
        for i, _ in enumerate(diff_arr):
            sum_diff += diff_arr[i]
            if nums[i] > sum_diff:
                return False
        return True
