from typing import List


class Solution:
    @staticmethod
    def numOfUnplacedFruits(fruits: List[int], baskets: List[int]) -> int:
        ans, baskets_length = len(fruits), len(baskets)
        flag = [False] * baskets_length
        for _, fruit in enumerate(fruits):
            for i, v in enumerate(baskets):
                if not flag[i] and fruit <= v:
                    flag[i] = True
                    ans -= 1
                    break
        return ans
