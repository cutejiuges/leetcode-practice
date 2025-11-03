from typing import List


class Solution:
    @staticmethod
    def minCost(colors: str, needed_time: List[int]) -> int:
        pre_color = colors[0]
        pre_time, cost = needed_time[0], 0
        for i in range(1, len(colors)):
            if pre_color == colors[i]:
                cost += min(pre_time, needed_time[i])
                pre_time = max(pre_time, needed_time[i])
            else:
                pre_color = colors[i]
                pre_time = needed_time[i]
        return cost
