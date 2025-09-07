from typing import List


class Solution:
    @staticmethod
    def sumZero(n: int) -> List[int]:
        ans = [0] * n
        low, high = 0, n - 1
        k = n
        while low < high:
            ans[low], ans[high] = -k, k
            low += 1
            high -= 1
            k -= 1
        if n & 1 == 1:
            ans[low] = 0
        return ans
