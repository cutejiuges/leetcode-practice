from typing import List


class Solution:
    @staticmethod
    def largestTriangleArea(points: List[List[int]]) -> float:
        max_area, n = 0, len(points)
        for i in range(n):
            for j in range(i+1, n):
                for k in range(j+1, n):
                    area = 0.5 * abs(
                        points[i][0] * (points[j][1] - points[k][1]) +
                        points[j][0] * (points[k][1] - points[i][1]) +
                        points[k][0] * (points[i][1] - points[j][1])
                    )
                    max_area = max(area, max_area)
        return max_area
