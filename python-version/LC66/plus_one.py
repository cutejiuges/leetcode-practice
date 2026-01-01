from typing import List


class Solution:
    @staticmethod
    def plusOne(digits: List[int]) -> List[int]:
        for i in range(len(digits)-1, -1, -1):
            digits[i] = (digits[i] + 1) % 10
            if digits[i] != 0:
                return digits

        ans = [0] * (len(digits) + 1)
        ans[0] = 1
        return ans
