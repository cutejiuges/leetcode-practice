from collections import defaultdict


class Solution:
    @staticmethod
    def numRabbits(answers: list[int]) -> int:
        ans = 0
        mp = defaultdict(int)
        for a in answers:
            if mp[a] == 0:
                ans += a + 1
                mp[a] = a
            else:
                mp[a] -= 1
        return ans
