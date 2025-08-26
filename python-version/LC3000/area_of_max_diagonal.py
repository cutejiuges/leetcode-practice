from typing import List


class Solution:
    @staticmethod
    def areaOfMaxDiagonal(dimensions: List[List[int]]) -> int:
        max_area, max_length = 0, 0
        for dimension in dimensions:
            area = dimension[0] * dimension[1]
            length = dimension[0]**2 + dimension[1]**2
            if length > max_length:
                max_length = length
                max_area = area
            elif length == max_length:
                max_area = max(max_area, area)
        return max_area
