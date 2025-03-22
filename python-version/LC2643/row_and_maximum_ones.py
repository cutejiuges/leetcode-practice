class Solution:
    def rowAndMaximumOnes(self, mat: list[list[int]]) -> list[int]:
        max_cnt_of_ones, index = -1, -1
        for i, row in enumerate(mat):
            ones_cnt = 0
            for m in row:
                ones_cnt += m
            if ones_cnt > max_cnt_of_ones:
                max_cnt_of_ones = ones_cnt
                index = i
        return [index, max_cnt_of_ones]
