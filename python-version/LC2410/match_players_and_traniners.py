from typing import List


class Solution:
    @staticmethod
    def matchPlayersAndTrainers(players: List[int], trainers: List[int]) -> int:
        players.sort()
        trainers.sort()
        m, n = len(players), len(trainers)
        i, j, ans = 0, 0, 0
        while i < m and j < n:
            if players[i] <= trainers[j]:
                ans += 1
                i += 1
            j += 1
        return ans
