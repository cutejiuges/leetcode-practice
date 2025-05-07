import heapq
from math import inf


class Pos:
    def __init__(self, x, y, dis):
        self.x = x
        self.y = y
        self.dis = dis

    def __lt__(self, other):
        return self.dis < other.dis


class Solution:
    @staticmethod
    def minTimeToReach(move_time: list[list[int]]) -> int:
        dirs = [(1, 0), (0, 1), (-1, 0), (0, -1)]
        row, col = len(move_time), len(move_time[0])
        distance = [[inf] * col for _ in range(row)]
        distance[0][0] = 0  # 起点是0,0 到起点的距离也是0

        min_heap = [Pos(0, 0, 0)]
        heapq.heapify(min_heap)
        while len(min_heap) > 0:
            pos = heapq.heappop(min_heap)
            if pos.x == row - 1 and pos.y == col - 1:
                return pos.dis
            if pos.dis > distance[pos.x][pos.y]:
                continue
            for d in dirs:
                x, y = pos.x + d[0], pos.y + d[1]
                if 0 <= x < row and 0 <= y < col:
                    new_dis = max(pos.dis, move_time[x][y]) + 1
                    if new_dis < distance[x][y]:
                        distance[x][y] = new_dis
                        heapq.heappush(min_heap, Pos(x, y, new_dis))
        return -1
