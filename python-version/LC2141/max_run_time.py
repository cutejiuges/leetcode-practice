from typing import List


class Solution:
    @staticmethod
    def maxRunTime(n: int, batteries: List[int]) -> int:
        total = sum(batteries)
        low, high = 0, total // n + 1
        while low + 1 < high:
            mid = low + (high - low) // 2
            temp = 0
            for num in batteries:
                temp += min(mid, num)
            if n * mid <= temp:
                low = mid
            else:
                high = mid
        return low
