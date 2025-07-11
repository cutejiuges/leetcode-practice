from typing import List


class Solution:
    @staticmethod
    def countDays(days: int, meetings: List[List[int]]) -> int:
        ans = days
        meetings.sort(key=lambda x: x[0])
        start, end = meetings[0][0], meetings[0][1]
        for i in range(1, len(meetings)):
            if meetings[i][0] > end:
                ans -= end - start + 1
                start = meetings[i][0]
            end = max(end, meetings[i][1])
        ans -= end - start + 1
        return ans
