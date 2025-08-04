from collections import defaultdict
from typing import List


class Solution:
    @staticmethod
    def totalFruit(fruits: List[int]) -> int:
        left, ans = 0, 0
        mp = defaultdict(int)
        for i, v in enumerate(fruits):
            mp[v] = mp.get(v, 0) + 1
            while len(mp) > 2:
                mp[fruits[left]] = mp[fruits[left]] - 1
                if mp[fruits[left]] == 0:
                    del mp[fruits[left]]
                left += 1
            ans = max(ans, i - left + 1)
        return ans
