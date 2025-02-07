class Solution:
    def generateMatrix(self, n: int) -> list[list[int]]:
        curNum, maxNum = 1, n*n
        dir = [[0, 1], [1, 0], [0, -1], [-1, 0]]
        curDir = 0
        matrix = [[0] * n for _ in range(n)]
        curRow, curCol = 0, 0
        while curNum <= maxNum:
            matrix[curRow][curCol] = curNum
            curNum += 1
            newRow, newCol = curRow + dir[curDir][0], curCol + dir[curDir][1]
            if newRow < 0 or newRow >= n or newCol < 0 or newCol >= n or matrix[newRow][newCol] != 0:
                curDir = (curDir + 1) % 4
            curRow, curCol = curRow + dir[curDir][0], curCol + dir[curDir][1]
        return matrix
    
