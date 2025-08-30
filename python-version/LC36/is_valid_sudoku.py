from typing import List


class Solution:
    @staticmethod
    def isValidSudoku(board: List[List[str]]) -> bool:
        row, col = [[False] * 9 for _ in range(9)], [[False] * 9 for _ in range(9)]
        box = [[[False] * 9 for _ in range(3)] for _ in range(3)]
        for i in range(9):
            for j in range(9):
                if board[i][j] == '.':
                    continue
                num = ord(board[i][j]) - ord('1')
                if row[i][num] or col[j][num] or box[i//3][j//3][num]:
                    return False
                row[i][num] = True
                col[j][num] = True
                box[i//3][j//3][num] = True
        return True
