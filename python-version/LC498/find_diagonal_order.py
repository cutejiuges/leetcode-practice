from typing import List


class Solution:
    @staticmethod
    def findDiagonalOrder(mat: List[List[int]]) -> List[int]:
        row, col = len(mat), len(mat[0])
        cnt = row + col - 1
        m, n, idx = 0, 0, 0
        res = [0] * (row * col)
        for i in range(cnt):
            if i & 1 == 0:
                while m >= 0 and n < col:
                    res[idx] = mat[m][n]
                    idx += 1
                    m -= 1
                    n += 1
                if n < col:
                    m += 1
                else:
                    m += 2
                    n -= 1
            else:
                while m < row and n >= 0:
                    res[idx] = mat[m][n]
                    idx += 1
                    m += 1
                    n -= 1
                if m < row:
                    n += 1
                else:
                    n += 2
                    m -= 1
        return res
