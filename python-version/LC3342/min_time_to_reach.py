import heapq
from math import inf


class Pos:
    def __init__(self, x: int, y: int, dist: int):
        self.x = x
        self.y = y
        self.dist = dist

    def __lt__(self, other):
        return self.dist < other.dist


class Solution:
    @staticmethod
    def minTimeToReach(move_time: list[list[int]]) -> int:
        dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        row, col = len(move_time), len(move_time[0])
        distance = [[inf] * col for _ in range(row)]

        min_heap = [Pos(0, 0, 0)]
        heapq.heapify(min_heap)
        while len(min_heap):
            poll = heapq.heappop(min_heap)
            if poll.x == row - 1 and poll.y == col - 1:
                return poll.dist
            if poll.dist > distance[poll.x][poll.y]:
                continue
            for d in dirs:
                new_x, new_y = poll.x + d[0], poll.y + d[1]
                time = 2 if (new_x + new_y) & 1 == 0 else 1
                if 0 <= new_x < row and 0 <= new_y < col:
                    new_dist = max(poll.dist, move_time[new_x][new_y]) + time
                    if new_dist < distance[new_x][new_y]:
                        distance[new_x][new_y] = new_dist
                        heapq.heappush(min_heap, Pos(new_x, new_y, new_dist))
        return -1
