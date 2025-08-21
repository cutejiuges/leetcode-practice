from typing import List


class Solution:
    def numSubmat(self, mat: List[List[int]]) -> int:
        if len(mat) == 0 or len(mat[0]) == 0:
            return 0
        m, n = len(mat), len(mat[0])
        temp = [[0] * n for _ in range(m)]
        ans = 0
        for i in range(m):
            for j in range(n):
                if j == 0:
                    temp[i][j] = mat[i][j]
                else:
                    temp[i][j] = 0 if mat[i][j] == 0 else temp[i][j - 1]+1
                cur = temp[i][j]
                for k in range(i, -1, -1):
                    cur = min(cur, temp[k][j])
                    if cur == 0:
                        break
                    ans += cur
        return ans
